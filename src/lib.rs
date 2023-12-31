use std::collections::HashMap;

use chrono::{offset::Utc, DateTime};
use rust_decimal::Decimal;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Default, Serialize, Deserialize)]
pub struct DetailedChannelSubscribe {
    pub name: String,
    pub product_ids: Vec<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ChannelSubscribe {
    NameOnly(String),
    Detailed(DetailedChannelSubscribe),
}

#[derive(Debug, Clone, Default, Serialize, Deserialize)]
pub struct Subscribe {
    pub product_ids: Vec<String>,
    pub channels: Vec<ChannelSubscribe>,
}

#[derive(Debug, Clone, Default, Serialize, Deserialize)]
pub struct Subscriptions {
    pub channels: Vec<ChannelSubscribe>,
}

#[derive(Debug, Clone, Default, Serialize, Deserialize)]
pub struct Unsubscribe {
    pub channels: Vec<ChannelSubscribe>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "order_type")]
pub enum ReceivedOrderType {
    #[serde(rename = "market")]
    Market { funds: Decimal },
    #[serde(rename = "limit")]
    Limit { size: Decimal, price: Decimal },
}

/// Received indicates an order received.
/// https://docs.cloud.coinbase.com/exchange/docs/websocket-channels#received
/// ```json
/// {
///   "type": "received",
///   "time": "2014-11-07T08:19:27.028459Z",
///   "product_id": "BTC-USD",
///   "sequence": 10,
///   "order_id": "d50ec984-77a8-460a-b958-66f114b0de9b",
///   "size": "1.34",
///   "price": "502.1",
///   "side": "buy",
///   "order_type": "limit",
///   "client-oid": "d50ec974-76a2-454b-66f135b1ea8c"
/// }
/// ```
/// note the `client-oid` may be empty.
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Received {
    pub time: DateTime<Utc>,
    pub product_id: String,
    pub sequence: u64,
    pub order_id: String,
    #[serde(flatten)]
    pub order_type: ReceivedOrderType,
    pub side: String,
    #[serde(rename = "client-oid")]
    pub client_oid: Option<String>,
}

/// Open indicates an order is open on the full channel.
/// https://docs.cloud.coinbase.com/exchange/docs/websocket-channels#open
/// ```json
/// {
///   "type": "open",
///   "time": "2014-11-07T08:19:27.028459Z",
///   "product_id": "BTC-USD",
///   "sequence": 10,
///   "order_id": "d50ec984-77a8-460a-b958-66f114b0de9b",
///   "price": "200.2",
///   "remaining_size": "1.00",
///   "side": "sell"
/// }
/// ```
#[derive(Debug, Clone, Serialize, Deserialize, Default)]
pub struct Open {
    pub time: DateTime<Utc>,
    pub product_id: String,
    pub sequence: u64,
    pub order_id: String,
    pub price: Decimal,
    pub remaining_size: Decimal,
    pub side: String,
}

/// Done indicates an order is done on the full channel.
/// ```json
/// {
///   "type": "done",
///   "time": "2014-11-07T08:19:27.028459Z",
///   "product_id": "BTC-USD",
///   "sequence": 10,
///   "price": "200.2",
///   "order_id": "d50ec984-77a8-460a-b958-66f114b0de9b",
///   "reason": "filled", // or "canceled"
///   "side": "sell",
///   "remaining_size": "0"
/// }
/// ```
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Done {
    pub time: DateTime<Utc>,
    pub product_id: String,
    pub sequence: u64,
    pub price: Option<Decimal>,
    pub order_id: String,
    pub reason: String,
    pub cancel_reason: Option<String>,
    pub side: String,
    pub remaining_size: Decimal,
}

/// Match is a match between two orders on the full channel.
/// ```json
/// {
///   "type": "match",
///   "trade_id": 10,
///   "sequence": 50,
///   "maker_order_id": "ac928c66-ca53-498f-9c13-a110027a60e8",
///   "taker_order_id": "132fb6ae-456b-4654-b4e0-d681ac05cea1",
///   "time": "2014-11-07T08:19:27.028459Z",
///   "product_id": "BTC-USD",
///   "size": "5.23512",
///   "price": "400.23",
///   "side": "sell"
/// }
/// ```
#[derive(Debug, Clone, Serialize, Deserialize, Default)]
pub struct Match {
    pub trade_id: u64,
    pub sequence: u64,
    pub maker_order_id: String,
    pub taker_order_id: String,
    pub time: DateTime<Utc>,
    pub product_id: String,
    pub size: Decimal,
    pub price: Decimal,
    pub side: String,
}

/// Change to the order from the full channel.
/// ```json
/// {
///   "type": "change",
///   "reason":"STP",
///   "time": "2014-11-07T08:19:27.028459Z",
///   "sequence": 80,
///   "order_id": "ac928c66-ca53-498f-9c13-a110027a60e8",
///   "side": "sell",
///   "product_id": "BTC-USD",
///   "old_size": "12.234412",
///   "new_size": "5.23512",
///   "price": "400.23"
/// }
/// ```
/// Or by user modify the order
/// ```json
/// {
///   "type": "change",
///   "reason":"modify_order",
///   "time": "2022-06-06T22:55:43.433114Z",
///   "sequence": 24753,
///   "order_id": "c3f16063-77b1-408f-a743-88b7bc20cdcd",
///   "side": "buy",
///   "product_id": "ETH-USD",
///   "old_size": "80",
///   "new_size": "80",
///   "old_price": "7",
///   "new_price": "6"
/// }
/// ```
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Change {
    pub reason: String,
    pub time: DateTime<Utc>,
    pub sequence: u64,
    pub order_id: String,
    pub side: String,
    pub product_id: String,
    pub old_size: Option<Decimal>,
    pub new_size: Option<Decimal>,
    pub size: Option<Decimal>,
    pub old_price: Option<Decimal>,
    pub new_price: Option<Decimal>,
    pub price: Option<Decimal>,
}

/// Active is the stop order activation on the full channel.
/// ```json
/// {
///   "type": "activate",
///   "product_id": "test-product",
///   "timestamp": "1483736448.299000",
///   "user_id": "12",
///   "profile_id": "30000727-d308-cf50-7b1c-c06deb1934fc",
///   "order_id": "7b52009b-64fd-0a2a-49e6-d8a939753077",
///   "stop_type": "entry",
///   "side": "buy",
///   "stop_price": "80",
///   "size": "2",
///   "funds": "50",
///   "private": true
/// }
/// ```
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Activate {
    pub product_id: String,
    pub timestamp: String,
    pub user_id: String,
    pub profile_id: String,
    pub stop_type: String,
    pub side: String,
    pub stop_price: Decimal,
    pub size: Decimal,
    pub funds: Decimal,
    pub private: bool,
}

/// Level3 Feed
/// Level3 feeds first sends schema, then string arrays
/// https://docs.cloud.coinbase.com/exchange/docs/websocket-channels#level3-channel
///
/// ```json
/// {
///   "type": "level3",
///   "schema": {
///     "change": [
///       "type",
///       "product_id",
///       "sequence",
///       "order_id",
///       "price",
///       "size",
///       "time"
///     ],
///     "done": [
///       "type",
///       "product_id",
///       "sequence",
///       "order_id",
///       "time"
///     ],
///     "match": [
///       "type",
///       "product_id",
///       "sequence",
///       "maker_order_id",
///       "taker_order_id",
///       "price",
///       "size",
///       "time"
///     ],
///     "noop": [
///       "type",
///       "product_id",
///       "sequence",
///       "time"
///     ],
///     "open": [
///       "type",
///       "product_id",
///       "sequence",
///       "order_id",
///       "side",
///       "price",
///       "size",
///       "time"
///     ]
///   }
/// }
/// ```
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Level3 {
    schema: HashMap<String, Vec<String>>,
}

/// Error message coming from the api.
/// ```json
/// {"type":"error","message":"Failed to subscribe","reason":"ETH-USD is not a valid product"}
/// ```
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Error {
    pub message: String,
    pub reason: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(tag = "type")]
pub enum CBMessageStruct {
    #[serde(rename = "subscribe")]
    Subscribe(Subscribe),
    #[serde(rename = "subscriptions")]
    Subscriptions(Subscriptions),
    #[serde(rename = "unsubscribe")]
    Unsubscribe(Unsubscribe),
    #[serde(rename = "received")]
    Received(Received),
    #[serde(rename = "open")]
    Open(Open),
    #[serde(rename = "done")]
    Done(Done),
    #[serde(rename = "match")]
    Match(Match),
    #[serde(rename = "change")]
    Change(Change),
    #[serde(rename = "activate")]
    Activate(Activate),
    #[serde(rename = "level3")]
    Level3(Level3),
    #[serde(rename = "error")]
    Error(Error),
}

#[allow(clippy::large_enum_variant)]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CBMessage {
    CBMessageStruct(CBMessageStruct),
    StrVec(Vec<String>),
}

#[cfg(test)]
mod tests {
    use super::*;
    use serde_json::from_str;
    use std::fs::read_to_string;

    #[test]
    fn test_deserialize() {
        let received = r#"{
  "type": "received",
  "time": "2014-11-07T08:19:27.028459Z",
  "product_id": "BTC-USD",
  "sequence": 10,
  "order_id": "d50ec984-77a8-460a-b958-66f114b0de9b",
  "size": "1.34",
  "price": "502.1",
  "side": "buy",
  "order_type": "limit",
  "client-oid": "d50ec974-76a2-454b-66f135b1ea8c"
}
"#;
        let r = from_str(received);
        println!("{:?}", r);
        assert!(r.is_ok());
        let r: CBMessageStruct = r.unwrap();
        if let CBMessageStruct::Received(r) = r {
            assert_eq!(
                r.client_oid,
                Some("d50ec974-76a2-454b-66f135b1ea8c".to_owned())
            );
        } else {
            panic!("failed to parse message");
        }
    }

    #[test]
    fn test_many_deserialize() {
        let test_file = concat!(env!("CARGO_MANIFEST_DIR"), "/testdata/messages.json");
        let contents = read_to_string(test_file).expect("failed to read the messages");
        for aline in contents.split('\n') {
            if aline.is_empty() {
                continue;
            }
            let r = from_str::<CBMessageStruct>(aline);
            if r.is_err() {
                panic!("failed to parse {}: {:?}", aline, r.err());
            }
        }
    }
}
