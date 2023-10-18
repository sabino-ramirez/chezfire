package auth

const (
	AuthURL   = "https://accounts.spotify.com/authorize"
	TokenURL  = "https://accounts.spotify.com/api/token"
	GrantType = "client_credentials"
)

// scopes passed in auth request define user granted permissions
const (
	ScopeImageUpload               = "ugc-image-upload"
	ScopePlaylistReadPrivate       = "playlist-read-private"
	ScopePlaylistModifyPublic      = "playlist-modify-public"
	ScopePlaylistModifyPrivate     = "playlist-modify-private"
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	ScopeUserFollowModify          = "user-follow-modify"
	ScopeUserFollowRead            = "user-follow-read"
	ScopeUserLibraryModify         = "user-library-modify"
	ScopeUserLibraryRead           = "user-library-read"
	ScopeUserReadPrivate           = "user-read-private"
	ScopeUserReadEmail             = "user-read-email"
	ScopeUserReadCurrentlyPlaying  = "user-read-currently-playing"
	ScopeUserReadPlaybackState     = "user-read-playback-state"
	ScopeUserModifyPlaybackState   = "user-modify-playback-state"
	ScopeUserReadRecentlyPlayed    = "user-read-recently-played"
	ScopeUserTopRead               = "user-top-read"
	ScopeStreaming                 = "streaming"
)
