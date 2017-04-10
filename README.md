# matterpoll-emoji

Polling feature for mattermost's custom slash command.

## Setup server

Clone this repository.
```
git clone https://github.com/kaakaa/matterpoll-emoji.git
cd matterpoll-emoji
```

Write configuration of `matterpoll-emoji` to config.json
```
{
	"host": "http://localhost:8505",  # Your Mattermost server
	"user": {
		"id": "bot",           # existiong account info of your Mattermost
		"password": "botbot"   # (It's recommended create bot account.)
	}
}
```

Setup `matterpoll-emoji` server.
```
glide install
go run main.go -p 8505
```

## Setup mattermost

Create a `Custom Slash Command` from Integration > Slash Commands > Add Slash Command.

* DisplayName - Arbitrary (ex. MatterPoll)
* Description - Arbitrary (ex. Polling feature by https://github.com/kaakaa/matterpoll-emoji)
* Command Trigger Word - `poll`
* Request URL - http://localhost:8505/poll
* Request Method - `POST`
* Others - optional

## Usage

Typing this on mattermost

```
/poll `What do you gys wanna grab for lunch?` :pizza: :sushi: :fried_shrimp: :spaghetti: :apple:
```

then posting poll comment

![screen_shot](https://raw.githubusercontent.com/kaakaa/matterpoll-emoji/master/matterpoll-emoji.png)

## License
* MIT
  * see [LICENSE](LICENSE)

