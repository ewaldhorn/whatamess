pub mod georesponse;

use std::net::SocketAddr;

use axum::{extract::Query, routing::get, Router};
use georesponse::{GeoResponse, LatLong};
use reqwest::StatusCode;
use serde::Deserialize;

#[derive(Deserialize)]
pub struct WeatherQuery {
    pub city: String,
}

// basic handler that responds with a static string
async fn index() -> &'static str {
    "Index"
}

async fn weather(Query(params): Query<WeatherQuery>) -> Result<String, StatusCode> {
    let lat_long = fetch_lat_long(&params.city)
        .await
        .map_err(|_| StatusCode::NOT_FOUND)?;
    Ok(format!(
        "{}: {}, {}",
        params.city, lat_long.latitude, lat_long.longitude
    ))
}

#[axum_macros::debug_handler]
async fn stats() -> &'static str {
    "Stats"
}

async fn fetch_lat_long(city: &str) -> Result<LatLong, Box<dyn std::error::Error>> {
    let endpoint = format!(
        "https://geocoding-api.open-meteo.com/v1/search?name={}&count=1&language=en&format=json",
        city
    );
    let response = reqwest::get(&endpoint).await?.json::<GeoResponse>().await?;
    response
        .results
        .get(0)
        .cloned()
        .ok_or("No results found".into())

    // or
    // match response.results.get(0) {
    //     Some(lat_long) => Ok(lat_long.clone()),
    //     None => Err("No results found".into()),
    // }
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(index))
        .route("/weather", get(weather))
        .route("/stats", get(stats));

    let addr = SocketAddr::from(([127, 0, 0, 1], 9000));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}
