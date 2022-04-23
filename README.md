# Pi-Stock DE
<img src="https://github.com/julianfbeck/pi-stock-de/blob/main/frontend/public/android-chrome-512x512.png" height="100">

> Monitor and get notified if raspberry pis are available in german stores

## Build Using
- Backend:
  - Golang (hexagonal pattern)
  - Colly
  - fiber
- Frontend:
  - React
  - Mantine
  - React-Query
  - Typescript  

## Monitored Stores
- Bechtle
- BerryBase
- BuyZero
- ELV
- Funk24
- OKDO
- Pishop.ch
- Rasppishop
- Semaf.at
- welectron

## Get Notified using
- Webhooks
- Email
- Pushover

## Planned Features:
- [x] Delete-Route to enable users to unsubscribe from Notifications
- [x] Add more notification Services
  - [ ] Web-Push-Notifications
  - [x] Add RSS-Feed
- [x] Protect Create-Notification Route with Captcha
- [x] New Custom Fav-Icon

## Installation
The easiest way to use PI-STOCK is using the Docker Image 

```
docker pull kickbeak/pi-stock-de
docker run -d -p 3001:3001 --env-file .env --name pi-stock-de kickbeak/pi-stock-de 
```
### Environment Variables
Supply the environment variables using the `.env` file.
See the `.env.example` file for all environment variables.

## Development
PI-STOCK-DE is build using golang, go-fiber and colly for scraping and React together with mantine for the frontend.
```
go run main.go
cd frontend
npm build
```
To develop the frontend live, you need to change the cors settings for the go-fiber config.

## Add new Shop
To add a new shop, you can simply create a new adaptors inside the `adaptors` folder.

Add the Shop URLs to monitor to the `websites.yaml` file, and initalize the adaptor inside the `main.go` file.

