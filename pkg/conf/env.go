package conf

// legacyEnvToFieldName maps from the legacy env var name to the
// SiteConfiguration struct field that the env var sets. Only top-level fields
// are traversed, so we needn't worry about nested names or collisions.
var legacyEnvToFieldName = map[string]string{
	"AdminUsernames": "ADMIN_USERNAMES",
	"AppID":          "TRACKING_APP_ID",
	"AppURL":         "SRC_APP_URL",
	// AuthUserOrgMap has no env var
	"AutoRepoAdd": "AUTO_REPO_ADD",
	"CorsOrigin":  "CORS_ORIGIN",
	// DisablePublicRepoRedirects is handled specially (inverted PUBLIC_REPO_REDIRECTS)
	"DisableTelemetry":               "DISABLE_TELEMETRY",
	"ExecuteGradleOriginalRootPaths": "EXECUTE_GRADLE_ORIGINAL_ROOT_PATHS",
	"GitOriginMap":                   "ORIGIN_MAP",
	"Github":                         "GITHUB_CONFIG",
	"GithubClientID":                 "GITHUB_CLIENT_ID",
	"GithubClientSecret":             "GITHUB_CLIENT_SECRET",
	"GitoliteHosts":                  "GITOLITE_HOSTS",
	"HtmlBodyBottom":                 "HTML_BODY_BOTTOM",
	"HtmlBodyTop":                    "HTML_BODY_TOP",
	"HtmlHeadBottom":                 "HTML_HEAD_BOTTOM",
	"HtmlHeadTop":                    "HTML_HEAD_TOP",
	"InactiveRepos":                  "INACTIVE_REPOS",
	"LicenseKey":                     "LICENSE_KEY",
	"LightstepAccessToken":           "LIGHTSTEP_ACCESS_TOKEN",
	"LightstepProject":               "LIGHTSTEP_PROJECT",
	"MandrillKey":                    "MANDRILL_KEY",
	"MaxReposToSearch":               "MAX_REPOS_TO_SEARCH",
	"NoGoGetDomains":                 "NO_GO_GET_DOMAINS",
	"OidcClientID":                   "OIDC_CLIENT_ID",
	"OidcClientSecret":               "OIDC_CLIENT_SECRET",
	"OidcEmailDomain":                "OIDC_EMAIL_DOMAIN",
	"OidcOverrideToken":              "OIDC_OVERRIDE_TOKEN",
	"OidcProvider":                   "OIDC_OP",
	"Phabricator":                    "PHABRICATOR_CONFIG",
	"PrivateArtifactRepoID":          "PRIVATE_ARTIFACT_REPO_ID",
	"PrivateArtifactRepoPassword":    "PRIVATE_ARTIFACT_REPO_PASSWORD",
	"PrivateArtifactRepoURL":         "PRIVATE_ARTIFACT_REPO_URL",
	"PrivateArtifactRepoUsername":    "PRIVATE_ARTIFACT_REPO_USERNAME",
	"RepoListUpdateInterval":         "REPO_LIST_UPDATE_INTERVAL",
	"ReposList":                      "REPOS_LIST",
	"SamlIDProviderMetadataURL":      "SAML_ID_PROVIDER_METADATA_URL",
	"SamlSPCert":                     "SAML_CERT",
	"SamlSPKey":                      "SAML_KEY",
	"SearchScopes":                   "SEARCH_SCOPES",
	// Settings has no env var
	"SsoUserHeader": "SSO_USER_HEADER",
	"TlsCert":       "TLS_CERT",
	"TlsKey":        "TLS_KEY",
}
