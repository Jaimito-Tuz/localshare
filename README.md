# LocalShare

Share your local development app with anyone instantly using Cloudflare Tunnel.

Made for CS students who need to demo their apps without deploying to a server. No port forwarding, no static IP — just run it and share the link.

## Requirements

- cloudflared installed on your machine
- A Cloudflare tunnel token (get one free at [one.dash.cloudflare.com](https://one.dash.cloudflare.com))

## Install cloudflared

```bash
# Add Cloudflare GPG key
sudo mkdir -p --mode=0755 /usr/share/keyrings
curl -fsSL https://pkg.cloudflare.com/cloudflare-public-v2.gpg | sudo tee /usr/share/keyrings/cloudflare-public-v2.gpg >/dev/null

# Add repo
echo 'deb [signed-by=/usr/share/keyrings/cloudflare-public-v2.gpg] https://pkg.cloudflare.com/cloudflared any main' | sudo tee /etc/apt/sources.list.d/cloudflared.list

# Install
sudo apt-get update && sudo apt-get install cloudflared
```

## Setup

1. Create a `.env` file in the same folder as the binary:

```
CLOUDFLARE_TUNNEL_TOKEN=your_token_here
```

2. Run localshare pointing to your app's port:

```bash
localshare --port 5000
```

## Download

Grab the prebuilt binary from [Releases](../../releases).

> Building from source requires cgo and native cloudflared headers which are a pain to set up. The prebuilt binary is the recommended path.

## Notes from the author

I built this for my networking class demo. Couldn't get the source to compile cleanly on my Linux VM because of the cgo deps, so I just ship a prebuilt binary. It worked first try for me — my professor accessed my app from his laptop with no issues.

## License

MIT
