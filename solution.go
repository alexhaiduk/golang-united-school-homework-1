package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	var Message string = emoji.Sprint("Hello :world_map:!")
	return Message
}
