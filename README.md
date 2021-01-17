# Channel Guard Plugin

Use this plugin to make channels read-only to some and writeable to other users. Channel Admins, Team admin, bots and system are all allowed to post. 

## Configuration

1. Go to **System Console > Plugins > Management** and click **Enable** to enable the Channel Guard plugin.

2. Modify your `config.json` file to include your Guards, under the `PluginSettings`. See below for an example of what this should look like.

## Usage

To configure the Guard, edit your `config.json` file with the following format:

```
"plugins": {
	"com.mattermost.channel-guard": {
	    "message": "This channel is under guard. You do not have the permissions to post. Please contact the system administrators if you believe this is incorrect"
		"guards": {
		    "channel-id-1": [ "user_1", "user_2" ],
		    "channel-id-2": [ "user_1", "user_2", "user_3" ],
		}
	}
}
```

where

- **channel-id-1**: Is the id of the channel you want to guard
- **user_1, user_2, etc**: List of Mattermost usernames that can post to the channel


## Development
For additional information on developing plugins, refer to [our plugin developer documentation](https://developers.mattermost.com/extend/plugins/).
