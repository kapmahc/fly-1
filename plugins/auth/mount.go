package auth

import "net/http"

// Mount mount web points
func (p *Plugin) Mount() {
	ug := p.Router.PathPrefix("/users").Subrouter().StrictSlash(true)
	ug.HandleFunc("/", p.Wrapper.JSON(p.indexUsers)).Methods(http.MethodGet)
	ug.HandleFunc("/sign-up", p.Wrapper.Form(&fmSignUp{}, p.postUsersSignUp)).Methods(http.MethodPost)
	ug.HandleFunc("/sign-in", p.Wrapper.Form(&fmSignIn{}, p.postUsersSignIn)).Methods(http.MethodPost)
	ug.HandleFunc("/confirm/{token}", p.Wrapper.Handle(p.getUsersConfirm)).Methods(http.MethodGet)
	ug.HandleFunc("/confirm", p.Wrapper.Form(&fmEmail{}, p.postUsersConfirm)).Methods(http.MethodPost)
	ug.HandleFunc("/unlock/{token}", p.Wrapper.Handle(p.getUsersUnlock)).Methods(http.MethodGet)
	ug.HandleFunc("/unlock", p.Wrapper.Form(&fmEmail{}, p.postUsersUnlock)).Methods(http.MethodPost)
	ug.HandleFunc("/forgot-password", p.Wrapper.Form(&fmEmail{}, p.postUsersForgotPassword)).Methods(http.MethodPost)
	ug.HandleFunc("/reset-password", p.Wrapper.Form(&fmResetPassword{}, p.postUsersResetPassword)).Methods(http.MethodPost)
	ug.Handle("/info", p.Wrapper.Wrap(p.Wrapper.JSON(p.getUsersInfo), p.Jwt.MustSignInMiddleware)).Methods(http.MethodGet)
	ug.Handle("/info", p.Wrapper.Wrap(p.Wrapper.Form(&fmInfo{}, p.postUsersInfo), p.Jwt.MustSignInMiddleware)).Methods(http.MethodPost)
	ug.Handle("/change-password", p.Wrapper.Wrap(p.Wrapper.Form(&fmChangePassword{}, p.postUsersChangePassword), p.Jwt.MustSignInMiddleware)).Methods(http.MethodGet)
	ug.Handle("/logs", p.Wrapper.Wrap(p.Wrapper.JSON(p.getUsersLogs), p.Jwt.MustSignInMiddleware)).Methods(http.MethodGet)
	ug.Handle("/sign-out", p.Wrapper.Wrap(p.Wrapper.JSON(p.deleteUsersSignOut), p.Jwt.MustSignInMiddleware)).Methods(http.MethodDelete)

	ag := p.Router.PathPrefix("/attachments").Subrouter().StrictSlash(true)
	ag.Handle("/", p.Wrapper.Wrap(p.Wrapper.JSON(p.indexAttachments), p.Jwt.MustSignInMiddleware)).Methods(http.MethodGet)
	ag.Handle("/", p.Wrapper.Wrap(p.Wrapper.Form(&fmAttachmentNew{}, p.createAttachment), p.Jwt.MustSignInMiddleware)).Methods(http.MethodPost)
	ag.Handle("/{id}", p.Wrapper.Wrap(p.Wrapper.JSON(p.showAttachment), p.Jwt.MustSignInMiddleware, p.canEditAttachment)).Methods(http.MethodGet)
	ag.Handle("/{id}", p.Wrapper.Wrap(p.Wrapper.Form(&fmAttachmentEdit{}, p.updateAttachment), p.Jwt.MustSignInMiddleware, p.canEditAttachment)).Methods(http.MethodPost)
	ag.Handle("/{id}", p.Wrapper.Wrap(p.Wrapper.JSON(p.destroyAttachment), p.Jwt.MustSignInMiddleware, p.canEditAttachment)).Methods(http.MethodDelete)
}
