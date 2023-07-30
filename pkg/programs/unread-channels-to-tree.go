package programs

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xlab/treeprint"
	v2 "traQ-Script/pkg/api/traQ/v2"
)

type channelTree struct {
	Id        string
	Name      string
	unreadNum int
	Children  []*channelTree
}

type channelTreeMapType map[string]*channelTree

func (c *channelTree) addChildren(children *channelTree) {
	c.Children = append(c.Children, children)
}

func (c *channelTree) string() string {
	if c.unreadNum > 0 {
		return fmt.Sprintf("%s (%d)", c.Name, c.unreadNum)
	} else {
		return c.Name
	}
}

func UnreadChannelsToTree() {
	unreadChannels, _, err := v2.GetMyUnreadChannels(v2.Me)
	if err != nil {
		panic(err)
	}

	channelTreeMap := channelTreeMapType{}
	rootChannel := &channelTree{
		Id:       "root",
		Name:     "root",
		Children: []*channelTree{},
	}
	channelTreeMap["root"] = rootChannel

	for _, unreadChannel := range unreadChannels {
		addThisChannelAndAllThisChannelParents(unreadChannel.ChannelId, channelTreeMap)
		channelTreeMap[unreadChannel.ChannelId].unreadNum = int(unreadChannel.Count)
	}

	v2.PostMessage(v2.Bot, v2.MyBotChannelId, ":@pikachu:の未読チャンネル一覧\n```\n"+channelTreeToTreeString(rootChannel)+"\n```", false)
}

func addThisChannelAndAllThisChannelParents(channelId string, channelTreeMap channelTreeMapType) {
	for {
		parentId, ok := v2.ChannelIdToChannelParentIdFunc(channelId)
		if !ok {
			addThisChannelToParentChildren(channelId, "root", channelTreeMap)
			break
		} else {
			addThisChannelToParentChildren(channelId, parentId, channelTreeMap)
			channelId = parentId
		}
	}
}

func addThisChannelToParentChildren(thisChannelId string, parentChannelId string, channelTreeMap channelTreeMapType) {
	makeNewChannelIfNotExist(parentChannelId, channelTreeMap)
	makeNewChannelIfNotExist(thisChannelId, channelTreeMap)
	childrenChannelTree := channelTreeMap[thisChannelId]
	parentChannelTree := channelTreeMap[parentChannelId]
	if !lo.Contains(parentChannelTree.Children, childrenChannelTree) {
		parentChannelTree.addChildren(childrenChannelTree)
	}
}

func makeNewChannelIfNotExist(channelId string, channelTreeMap channelTreeMapType) {
	_, exist := channelTreeMap[channelId]
	if !exist {
		newChannel := &channelTree{
			Id:       channelId,
			Name:     v2.ChannelIdToChannelNameFunc(channelId),
			Children: []*channelTree{},
		}
		channelTreeMap[channelId] = newChannel
	}
}

func channelTreeToTreeString(rootChannelTree *channelTree) string {
	tree := treeprint.New()

	for _, root := range rootChannelTree.Children {
		rootBranch := tree.AddBranch(root.string())
		addBranch(rootBranch, root.Children)
	}

	return tree.String()
}

func addBranch(parent treeprint.Tree, children []*channelTree) {
	for _, child := range children {
		branch := parent.AddBranch(child.string())
		addBranch(branch, child.Children)
	}
}
