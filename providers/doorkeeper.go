package providers

import (
	"net/url"
)

type DoorkeeperProvider struct {
	*ProviderData
}

func NewDoorkeeperProvider(p *ProviderData) *DoorkeeperProvider {
	p.ProviderName = "Alveo"

	DoorkeeperHost := p.OAuthProviderHost

	if p.LoginURL == nil || p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   DoorkeeperHost,
			Path:   "/oauth/authorize",
		}
	}
	if p.RedeemURL == nil || p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   DoorkeeperHost,
			Path:   "/oauth/token",
		}
	}
	// ValidationURL is the API Base URL
	if p.ValidateURL == nil || p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   DoorkeeperHost,
			Path:   "/oauth/authorize",
		}
	}
	//if p.Scope == "" {
	//	p.Scope = "user:email"
	//}
	return &DoorkeeperProvider{ProviderData: p}
}

func (p *DoorkeeperProvider) GetEmailAddress(s *SessionState) (string, error) {
	return "support@alveo.edu.au", nil
}

func (p *DoorkeeperProvider) GetUserName(s *SessionState) (string, error) {
	return "alveo.hcsvlab", nil
}
