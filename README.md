## Project Setup

This application uses **ProtonVPN** and **WireGuard** to securely send Telegram messages within a **Docker** environment. Follow the steps below to set up the project.

### Prerequisites

* **Docker** and **Docker Compose** installed on your system
* A valid **ProtonVPN** account
* Your own **WireGuard** configuration and **private keys**

---

### Setup Instructions

1. **Clone the repository**

   ```bash
   git clone https://github.com/salehrashid/Practice_go-telegram.git
   cd go-telegram
   ```

2. **Copy and rename the environment file**
   Start by copying the sample environment file:

   ```bash
   cp .env.sample .env
   ```

3. **Update the environment variables**
   Open the newly created `.env` file and **replace all placeholder values** (such as `PRIVATE_KEY_PLACEHOLDER`) with your **actual keys**:

   ```
   TELEGRAM_BOT_TOKEN=your-telegram-token
   WG_PRIVATE_KEY=your-actual-private-key
   ```

4. **Build and run the Docker containers**

   ```bash
   docker compose up --build
   ```

5. **Verify the application**
   The app will automatically connect to **ProtonVPN** via **WireGuard** and start sending Telegram messages securely.

---

### Tech Stack

* **Golang**: Core backend application
* **WireGuard**: VPN tunneling
* **ProtonVPN**: VPN provider
* **Docker & Docker Compose**: Containerization and orchestration
