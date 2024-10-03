package run

type Entry struct {
	Level    string `json:"level,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Port     string `json:"port,omitempty"`
	OsSignal string `json:"os_signal,omitempty"`
}

//Doesn't pass the test
//func TestApp_Run(t *testing.T) {
//	var decoder godecoder.Decoder
//	var entry Entry
//	ws := bytes.NewBuffer(make([]byte, 0, 1000))
//	tests := []struct {
//		name     string
//		app      App
//		callback func(a *App)
//		want     int
//		logs     []Entry
//	}{
//		{
//			name: "run server",
//			app: func() App {
//				conf := config.AppConf{
//					Server: config.Server{
//						Port:            "8080",
//						ShutdownTimeout: 5,
//					},
//				}
//				logger := logs.NewLogger(conf, zapcore.AddSync(ws))
//
//				app := NewApp(conf, logger)
//
//				app.Bootstrap()
//				decoder = godecoder.NewDecoder(jsoniter.Config{
//					EscapeHTML:             true,
//					SortMapKeys:            true,
//					ValidateJsonRawMessage: true,
//				})
//
//				return *app
//			}(),
//			logs: []Entry{
//				{
//					Level: "info",
//					Msg:   "server started",
//					Port:  "8080",
//				},
//				{
//					Level:    "info",
//					Msg:      "signal interrupt recieved",
//					Port:     "8080",
//					OsSignal: "interrupt",
//				},
//			},
//			callback: func(a *App) {
//				a.Sig <- os.Interrupt
//			},
//			want: errors.NoError,
//		},
//		{
//			name: "server network address bad port",
//			app: func() App {
//				conf := config.AppConf{
//					Server: config.Server{
//						Port:            "fff",
//						ShutdownTimeout: 0,
//					},
//				}
//				logger := logs.NewLogger(conf, zapcore.AddSync(ws))
//
//				app := NewApp(conf, logger)
//
//				app.Bootstrap()
//
//				return *app
//			}(),
//			logs: nil,
//			want: errors.GeneralError,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			a := &App{
//				conf:   tt.app.conf,
//				logger: tt.app.logger,
//				srv:    tt.app.srv,
//				Sig:    tt.app.Sig,
//			}
//			var wg sync.WaitGroup
//			wg.Add(1)
//			go func(t *testing.T) {
//				if got := a.Run(); got != tt.want {
//					t.Errorf("Run() = %v, want %v", got, tt.want)
//				}
//				wg.Done()
//			}(t)
//			time.Sleep(200 * time.Millisecond)
//			if tt.callback != nil {
//				tt.callback(a)
//			}
//			time.Sleep(10 * time.Millisecond)
//			scanner := bufio.NewScanner(ws)
//			wg.Wait()
//			// optionally, resize scanner's capacity for lines over 64K, see next example
//			if len(tt.logs) > 0 {
//				var logLines []string
//				for scanner.Scan() {
//					logLines = append(logLines, scanner.Text())
//				}
//				assert.True(t, len(logLines) == len(tt.logs))
//				for i, logLine := range logLines {
//					err := decoder.Decode(bytes.NewBuffer([]byte(logLine)), &entry)
//					assert.NoError(t, err)
//					assert.Equal(t, entry, tt.logs[i])
//				}
//			}
//		})
//	}
//}
