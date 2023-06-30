use cbrs::*;
use serde_json::{from_str, to_string_pretty};
use tungstenite::{connect, Message};
use url::Url;

const CB_WS_ENDPOINT: &str = "wss://ws-feed-public.sandbox.exchange.coinbase.com";

fn main() {
    let sub = Subscribe {
        channels: vec![ChannelSubscribe::Detailed(DetailedChannelSubscribe {
            name: "full".to_owned(),
            product_ids: vec!["ETH-BTC".to_owned(), "ETH-USD".to_owned()],
        })],
        product_ids: vec!["ETH-USD".to_owned(), "BTC-USD".to_owned()],
    };

    let sub = CBMesasge::Subscribe(sub);
    println!("{}", to_string_pretty(&sub).unwrap());

    let url =
        Url::parse(CB_WS_ENDPOINT).unwrap_or_else(|_| panic!("failed to parse {CB_WS_ENDPOINT}"));

    let (mut socket, x) =
        connect(url).unwrap_or_else(|_| panic!("failed to connect to {CB_WS_ENDPOINT}"));
    if x.status().as_u16() >= 400 {
        panic!("connecting to {CB_WS_ENDPOINT} unsuccessful, {x:?}");
    }

    let sub =
        to_string_pretty(&sub).unwrap_or_else(|_| panic!("failed to marshal {sub:?} to json."));

    let r = socket.write_message(tungstenite::Message::Text(sub));

    if r.is_err() {
        panic!("failed to send sub message: {:?}", r.err());
    }

    loop {
        let msg = socket.read_message();
        if msg.is_err() {
            panic!("reading error: {:?}", msg.err());
        }

        if let Message::Text(msg) = msg.unwrap() {
            let parsed = from_str::<CBMesasge>(&msg);
            if parsed.is_err() {
                println!("failed to parse msg: {}", msg);
                println!("failed reason: {:?}", parsed.err());
            } else {
                println!("{:?}", parsed);
            }
        } else {
            continue;
        }
    }
}
