use std::{fs::File, io::Write};

use cbrs::*;

use chrono::Utc;
use clap::Parser;
use futures_util::{SinkExt, StreamExt};
use serde_json::{from_str, to_string_pretty};
use tokio::signal::unix::{signal, SignalKind};
use tokio_tungstenite::{connect_async, tungstenite::Message};
use url::Url;

#[derive(Parser)]
struct Args {
    #[clap(long, short, default_value = "stat.csv")]
    output: String,
    #[clap(long, short, default_value = "full")]
    channel: String,
    #[clap(
        long,
        short,
        default_value = "wss://ws-feed-public.sandbox.exchange.coinbase.com"
    )]
    endpoint: String,
}

fn get_now() -> i64 {
    Utc::now().timestamp_nanos()
}

#[tokio::main]
async fn main() {
    let Args {
        output,
        channel,
        endpoint,
    } = Args::parse();

    let sub = Subscribe {
        channels: vec![ChannelSubscribe::Detailed(DetailedChannelSubscribe {
            name: channel.clone(),
            product_ids: vec!["BTC-USD".to_owned()],
        })],
        product_ids: vec![],
    };

    let sub = CBMesasgeStruct::Subscribe(sub);
    println!("{}", to_string_pretty(&sub).unwrap());

    let url = Url::parse(&endpoint).unwrap_or_else(|_| panic!("failed to parse {endpoint}"));

    let (socket, x) = connect_async(url)
        .await
        .unwrap_or_else(|_| panic!("failed to connect to {endpoint}"));

    if x.status().as_u16() >= 400 {
        panic!("connecting to {endpoint} unsuccessful, {x:?}");
    }

    let sub =
        to_string_pretty(&sub).unwrap_or_else(|_| panic!("failed to marshal {sub:?} to json."));

    let (mut sink, mut read) = socket.split();

    let r = sink.send(Message::Text(sub)).await;

    if r.is_err() {
        panic!("failed to send sub message: {:?}", r.err());
    }

    tokio::spawn(async move {
        let mut sigint = signal(SignalKind::interrupt()).unwrap();

        let _ = sigint.recv().await;

        let unsub = to_string_pretty(&CBMesasgeStruct::Unsubscribe(Unsubscribe {
            channels: vec![ChannelSubscribe::Detailed(DetailedChannelSubscribe {
                name: channel,
                product_ids: vec!["BTC-USD".to_owned()],
            })],
        }))
        .unwrap();
        println!("sending unsub");
        if sink.send(Message::Text(unsub)).await.is_ok() {
            let _ = sink.send(Message::Close(None)).await;
        }

        let _ = sink.close().await;
    });

    let mut stat_file = File::create(output).expect("failed to create stat.csv");
    let mut count: usize = 0;
    let mut size: usize = 0;
    writeln!(stat_file, "timestamp,count,size").expect("failed to write to stat");
    writeln!(stat_file, "{:?},{},{}", get_now(), count, size).expect("failed to write to stat");
    loop {
        let msg = read.next().await;
        if msg.is_none() {
            break;
        }
        let msg = msg.unwrap();
        if msg.is_err() {
            panic!("reading error: {:?}", msg.err());
        }

        if let Message::Text(msg) = msg.unwrap() {
            count += 1;
            size += msg.len();

            if count % 10_000 == 0 {
                writeln!(stat_file, "{:?},{},{}", get_now(), count, size)
                    .expect("failed to write to stat");
            }

            let parsed = from_str::<CBMessage>(&msg);
            if parsed.is_err() {
                println!("failed to parse msg: {}", msg);
                println!("failed reason: {:?}", parsed.err());
            } else {
                println!("{:?}", parsed.unwrap());
            }
        } else {
            continue;
        }
    }

    writeln!(stat_file, "{:?},{},{}", get_now(), count, size).expect("failed to write to stat");
}
