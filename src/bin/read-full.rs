const CB_WS_ENDPOINT: &str = "wss://ws-feed-public.sandbox.exchange.coinbase.com";

use serde_json::to_string_pretty;

use cbrs::*;

fn main() {
    let sub = Subscribe {
        channels: vec![
            ChannelSubscribe::NameOnly("level2".to_owned()),
            ChannelSubscribe::NameOnly("heartbeat".to_owned()),
            ChannelSubscribe::Detailed(DetailedChannelSubscribe {
                name: "ticker".to_owned(),
                product_ids: vec!["ETH-BTC".to_owned(), "ETH-USD".to_owned()],
            }),
        ],
        product_ids: vec!["ETH-USD".to_owned(), "BTC-USD".to_owned()],
    };

    println!("{}", to_string_pretty(&CBMesasge::Subscribe(sub)).unwrap());
}
