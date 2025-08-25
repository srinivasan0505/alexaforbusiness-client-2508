/**
 * MCP Server function for Creates a network profile with the specified details.
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("X-Amz-Target")) {
            queryParams.add("X-Amz-Target=" + args.get("X-Amz-Target"));
        }
        if (args.containsKey("Description")) {
            queryParams.add("Description=" + args.get("Description"));
        }
        if (args.containsKey("NetworkProfileName")) {
            queryParams.add("NetworkProfileName=" + args.get("NetworkProfileName"));
        }
        if (args.containsKey("SecurityType")) {
            queryParams.add("SecurityType=" + args.get("SecurityType"));
        }
        if (args.containsKey("ClientRequestToken")) {
            queryParams.add("ClientRequestToken=" + args.get("ClientRequestToken"));
        }
        if (args.containsKey("EapMethod")) {
            queryParams.add("EapMethod=" + args.get("EapMethod"));
        }
        if (args.containsKey("Ssid")) {
            queryParams.add("Ssid=" + args.get("Ssid"));
        }
        if (args.containsKey("CertificateAuthorityArn")) {
            queryParams.add("CertificateAuthorityArn=" + args.get("CertificateAuthorityArn"));
        }
        if (args.containsKey("NextPassword")) {
            queryParams.add("NextPassword=" + args.get("NextPassword"));
        }
        if (args.containsKey("Tags")) {
            queryParams.add("Tags=" + args.get("Tags"));
        }
        if (args.containsKey("TrustAnchors")) {
            queryParams.add("TrustAnchors=" + args.get("TrustAnchors"));
        }
        if (args.containsKey("CurrentPassword")) {
            queryParams.add("CurrentPassword=" + args.get("CurrentPassword"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_alexa_for_business_create_network_profile" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> X-Amz-TargetProperty = new HashMap<>();
        X-Amz-TargetProperty.put("type", "string");
        X-Amz-TargetProperty.put("required", true);
        X-Amz-TargetProperty.put("description", "");
        properties.put("X-Amz-Target", X-Amz-TargetProperty);
        Map<String, Object> DescriptionProperty = new HashMap<>();
        DescriptionProperty.put("type", "string");
        DescriptionProperty.put("required", false);
        DescriptionProperty.put("description", "");
        properties.put("Description", DescriptionProperty);
        Map<String, Object> NetworkProfileNameProperty = new HashMap<>();
        NetworkProfileNameProperty.put("type", "string");
        NetworkProfileNameProperty.put("required", true);
        NetworkProfileNameProperty.put("description", "");
        properties.put("NetworkProfileName", NetworkProfileNameProperty);
        Map<String, Object> SecurityTypeProperty = new HashMap<>();
        SecurityTypeProperty.put("type", "string");
        SecurityTypeProperty.put("required", true);
        SecurityTypeProperty.put("description", "");
        properties.put("SecurityType", SecurityTypeProperty);
        Map<String, Object> ClientRequestTokenProperty = new HashMap<>();
        ClientRequestTokenProperty.put("type", "string");
        ClientRequestTokenProperty.put("required", true);
        ClientRequestTokenProperty.put("description", "Input parameter: A unique, user-specified identifier for the request that ensures idempotency.");
        properties.put("ClientRequestToken", ClientRequestTokenProperty);
        Map<String, Object> EapMethodProperty = new HashMap<>();
        EapMethodProperty.put("type", "string");
        EapMethodProperty.put("required", false);
        EapMethodProperty.put("description", "");
        properties.put("EapMethod", EapMethodProperty);
        Map<String, Object> SsidProperty = new HashMap<>();
        SsidProperty.put("type", "string");
        SsidProperty.put("required", true);
        SsidProperty.put("description", "");
        properties.put("Ssid", SsidProperty);
        Map<String, Object> CertificateAuthorityArnProperty = new HashMap<>();
        CertificateAuthorityArnProperty.put("type", "string");
        CertificateAuthorityArnProperty.put("required", false);
        CertificateAuthorityArnProperty.put("description", "");
        properties.put("CertificateAuthorityArn", CertificateAuthorityArnProperty);
        Map<String, Object> NextPasswordProperty = new HashMap<>();
        NextPasswordProperty.put("type", "string");
        NextPasswordProperty.put("required", false);
        NextPasswordProperty.put("description", "");
        properties.put("NextPassword", NextPasswordProperty);
        Map<String, Object> TagsProperty = new HashMap<>();
        TagsProperty.put("type", "string");
        TagsProperty.put("required", false);
        TagsProperty.put("description", "");
        properties.put("Tags", TagsProperty);
        Map<String, Object> TrustAnchorsProperty = new HashMap<>();
        TrustAnchorsProperty.put("type", "string");
        TrustAnchorsProperty.put("required", false);
        TrustAnchorsProperty.put("description", "");
        properties.put("TrustAnchors", TrustAnchorsProperty);
        Map<String, Object> CurrentPasswordProperty = new HashMap<>();
        CurrentPasswordProperty.put("type", "string");
        CurrentPasswordProperty.put("required", false);
        CurrentPasswordProperty.put("description", "");
        properties.put("CurrentPassword", CurrentPasswordProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_alexa_for_business_create_network_profile",
            "Creates a network profile with the specified details.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileHandler(config));
    }
    
}