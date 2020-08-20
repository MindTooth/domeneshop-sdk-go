/*
 * Domeneshop API Documentation
 *
 * # Overview Domeneshop offers a simple, REST-based API, which currently supports the following  features: - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including: - Web hosting administration - Email address and email user/account administration  # Authentication The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website: <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package domeneshop
// SrvAllOf struct for SrvAllOf
type SrvAllOf struct {
	Type string `json:"type"`
	// The target hostname
	Data string `json:"data"`
	// SRV record priority, also known as preference. Lower values are usually preferred first
	Priority int32 `json:"priority"`
	// SRV record weight. Relevant if multiple records have same preference
	Weight int32 `json:"weight"`
	// SRV record port. The port where the service is found.
	Port int32 `json:"port"`
}
