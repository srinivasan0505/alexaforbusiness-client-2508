/**
 * MCP Server function for Creates a recurring schedule for usage reports to deliver to the specified S3 location with a specified daily or weekly interval.
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

class Post_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("S3BucketName")) {
            queryParams.add("S3BucketName=" + args.get("S3BucketName"));
        }
        if (args.containsKey("S3KeyPrefix")) {
            queryParams.add("S3KeyPrefix=" + args.get("S3KeyPrefix"));
        }
        if (args.containsKey("ScheduleName")) {
            queryParams.add("ScheduleName=" + args.get("ScheduleName"));
        }
        if (args.containsKey("Tags")) {
            queryParams.add("Tags=" + args.get("Tags"));
        }
        if (args.containsKey("ClientRequestToken")) {
            queryParams.add("ClientRequestToken=" + args.get("ClientRequestToken"));
        }
        if (args.containsKey("Format")) {
            queryParams.add("Format=" + args.get("Format"));
        }
        if (args.containsKey("Recurrence")) {
            queryParams.add("Recurrence=" + args.get("Recurrence"));
        }
        if (args.containsKey("ContentRange")) {
            queryParams.add("ContentRange=" + args.get("ContentRange"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_alexa_for_business_create_business_report_schedule" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleTool(MCPServer.APIConfig config) {
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
        Map<String, Object> S3BucketNameProperty = new HashMap<>();
        S3BucketNameProperty.put("type", "string");
        S3BucketNameProperty.put("required", false);
        S3BucketNameProperty.put("description", "");
        properties.put("S3BucketName", S3BucketNameProperty);
        Map<String, Object> S3KeyPrefixProperty = new HashMap<>();
        S3KeyPrefixProperty.put("type", "string");
        S3KeyPrefixProperty.put("required", false);
        S3KeyPrefixProperty.put("description", "");
        properties.put("S3KeyPrefix", S3KeyPrefixProperty);
        Map<String, Object> ScheduleNameProperty = new HashMap<>();
        ScheduleNameProperty.put("type", "string");
        ScheduleNameProperty.put("required", false);
        ScheduleNameProperty.put("description", "");
        properties.put("ScheduleName", ScheduleNameProperty);
        Map<String, Object> TagsProperty = new HashMap<>();
        TagsProperty.put("type", "string");
        TagsProperty.put("required", false);
        TagsProperty.put("description", "");
        properties.put("Tags", TagsProperty);
        Map<String, Object> ClientRequestTokenProperty = new HashMap<>();
        ClientRequestTokenProperty.put("type", "string");
        ClientRequestTokenProperty.put("required", false);
        ClientRequestTokenProperty.put("description", "");
        properties.put("ClientRequestToken", ClientRequestTokenProperty);
        Map<String, Object> FormatProperty = new HashMap<>();
        FormatProperty.put("type", "string");
        FormatProperty.put("required", true);
        FormatProperty.put("description", "");
        properties.put("Format", FormatProperty);
        Map<String, Object> RecurrenceProperty = new HashMap<>();
        RecurrenceProperty.put("type", "string");
        RecurrenceProperty.put("required", false);
        RecurrenceProperty.put("description", "");
        properties.put("Recurrence", RecurrenceProperty);
        Map<String, Object> ContentRangeProperty = new HashMap<>();
        ContentRangeProperty.put("type", "string");
        ContentRangeProperty.put("required", true);
        ContentRangeProperty.put("description", "");
        properties.put("ContentRange", ContentRangeProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_alexa_for_business_create_business_report_schedule",
            "Creates a recurring schedule for usage reports to deliver to the specified S3 location with a specified daily or weekly interval.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleHandler(config));
    }
    
}