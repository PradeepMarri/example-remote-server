package main

import (
	"github.com/mcp-server-everything-api/mcp-server/config"
	"github.com/mcp-server-everything-api/mcp-server/models"
	tools_message "github.com/mcp-server-everything-api/mcp-server/tools/message"
	tools_sse "github.com/mcp-server-everything-api/mcp-server/tools/sse"
	tools_token_exchange_key_prefixsha256_authorizationcode_if_data_returnundefined_constdecoded_decryptstring_encryptedtextdata "github.com/mcp-server-everything-api/mcp-server/tools/token_exchange_key_prefixsha256_authorizationcode_if_data_returnundefined_constdecoded_decryptstring_encryptedtextdata"
	tools_prefixsha256_key_if_data_returnundefined_constdecoded_decryptstring_encryptedtextdata "github.com/mcp-server-everything-api/mcp-server/tools/prefixsha256_key_if_data_returnundefined_constdecoded_decryptstring_encryptedtextdata"
	tools_fakeupstreamauth "github.com/mcp-server-everything-api/mcp-server/tools/fakeupstreamauth"
	tools_mcp "github.com/mcp-server-everything-api/mcp-server/tools/mcp"
	tools_mcp_logo_png "github.com/mcp-server-everything-api/mcp-server/tools/mcp_logo_png"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_message.CreatePost_messageTool(cfg),
		tools_sse.CreateGet_sseTool(cfg),
		tools_token_exchange_key_prefixsha256_authorizationcode_if_data_returnundefined_constdecoded_decryptstring_encryptedtextdata.CreateGet_token_exchange_key_prefixsha256authorizationcodeifdatareturnundefinedconstdecodeddecryptstringencryptedtextdataTool(cfg),
		tools_prefixsha256_key_if_data_returnundefined_constdecoded_decryptstring_encryptedtextdata.CreateGet_prefixsha256keyifdatareturnundefinedconstdecodeddecryptstringencryptedtextdataTool(cfg),
		tools_fakeupstreamauth.CreateGet_fakeupstreamauth_authorizeTool(cfg),
		tools_fakeupstreamauth.CreateGet_fakeupstreamauth_callbackTool(cfg),
		tools_mcp.CreateDelete_mcpTool(cfg),
		tools_mcp.CreateGet_mcpTool(cfg),
		tools_mcp.CreatePost_mcpTool(cfg),
		tools_mcp_logo_png.CreateGet_mcp_logo_pngTool(cfg),
	}
}
