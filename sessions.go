package swf

type Sessions struct {
	repository  map[string]*Session
}

type Session struct {
	sid		string
	is_auth bool
}

func (sessions *Sessions) Start(ctx *Context, username string) {
	session := &Session{
		sid: username,  // session_id要随机生成无规律字符串
		is_auth: true,
	}
	sessions.repository[username] = session

	// TODO: 使用ctx存储session, 最终在路由层返回设置HTTP头返回给客户端
}

func (sessions *Sessions)Set(sid string, is_auth bool) {
	sessions.repository[sid].is_auth = is_auth
}

func NewSessions() *Sessions {
	return &Sessions{
		repository: make(map[string]*Session, 0),
	}
}