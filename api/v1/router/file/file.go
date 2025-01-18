package file

func (r *Router) SetupFile() {
	file := r.Group("/api/file")
	{
		file.GET("/auth-pwd", r.AuthPwd)
		file.GET("/dir-list", r.Dirlist)
		file.GET("/download", r.Download)
		file.GET("/stream-video", r.StreamVideo)
		file.GET("/user-config", r.UserConfig)
	}
}
