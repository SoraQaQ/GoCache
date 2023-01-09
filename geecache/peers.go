package geecache

//根据传入的key选择相应节点
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

//从HTTP客户端中选择
type PeerGetter interface {
	//从对应的group查找缓存值
	Get(group string, key string) ([]byte, error)
}
