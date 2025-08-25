/**
 * MCP Server function for Updates an existing room profile by room profile ARN.
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

class Post_X_Amz_Target_Alexa_For_Business_Update_ProfileMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Alexa_For_Business_Update_ProfileHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("PSTNEnabled")) {
            queryParams.add("PSTNEnabled=" + args.get("PSTNEnabled"));
        }
        if (args.containsKey("WakeWord")) {
            queryParams.add("WakeWord=" + args.get("WakeWord"));
        }
        if (args.containsKey("DataRetentionOptIn")) {
            queryParams.add("DataRetentionOptIn=" + args.get("DataRetentionOptIn"));
        }
        if (args.containsKey("IsDefault")) {
            queryParams.add("IsDefault=" + args.get("IsDefault"));
        }
        if (args.containsKey("ProfileArn")) {
            queryParams.add("ProfileArn=" + args.get("ProfileArn"));
        }
        if (args.containsKey("ProfileName")) {
            queryParams.add("ProfileName=" + args.get("ProfileName"));
        }
        if (args.containsKey("TemperatureUnit")) {
            queryParams.add("TemperatureUnit=" + args.get("TemperatureUnit"));
        }
        if (args.containsKey("DistanceUnit")) {
            queryParams.add("DistanceUnit=" + args.get("DistanceUnit"));
        }
        if (args.containsKey("Timezone")) {
            queryParams.add("Timezone=" + args.get("Timezone"));
        }
        if (args.containsKey("Address")) {
            queryParams.add("Address=" + args.get("Address"));
        }
        if (args.containsKey("SetupModeDisabled")) {
            queryParams.add("SetupModeDisabled=" + args.get("SetupModeDisabled"));
        }
        if (args.containsKey("Locale")) {
            queryParams.add("Locale=" + args.get("Locale"));
        }
        if (args.containsKey("MaxVolumeLimit")) {
            queryParams.add("MaxVolumeLimit=" + args.get("MaxVolumeLimit"));
        }
        if (args.containsKey("MeetingRoomConfiguration")) {
            queryParams.add("MeetingRoomConfiguration=" + args.get("MeetingRoomConfiguration"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_alexa_for_business_update_profile" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Alexa_For_Business_Update_ProfileTool(MCPServer.APIConfig config) {
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
        Map<String, Object> PSTNEnabledProperty = new HashMap<>();
        PSTNEnabledProperty.put("type", "string");
        PSTNEnabledProperty.put("required", false);
        PSTNEnabledProperty.put("description", "");
        properties.put("PSTNEnabled", PSTNEnabledProperty);
        Map<String, Object> WakeWordProperty = new HashMap<>();
        WakeWordProperty.put("type", "string");
        WakeWordProperty.put("required", false);
        WakeWordProperty.put("description", "");
        properties.put("WakeWord", WakeWordProperty);
        Map<String, Object> DataRetentionOptInProperty = new HashMap<>();
        DataRetentionOptInProperty.put("type", "string");
        DataRetentionOptInProperty.put("required", false);
        DataRetentionOptInProperty.put("description", "");
        properties.put("DataRetentionOptIn", DataRetentionOptInProperty);
        Map<String, Object> IsDefaultProperty = new HashMap<>();
        IsDefaultProperty.put("type", "string");
        IsDefaultProperty.put("required", false);
        IsDefaultProperty.put("description", "");
        properties.put("IsDefault", IsDefaultProperty);
        Map<String, Object> ProfileArnProperty = new HashMap<>();
        ProfileArnProperty.put("type", "string");
        ProfileArnProperty.put("required", false);
        ProfileArnProperty.put("description", "");
        properties.put("ProfileArn", ProfileArnProperty);
        Map<String, Object> ProfileNameProperty = new HashMap<>();
        ProfileNameProperty.put("type", "string");
        ProfileNameProperty.put("required", false);
        ProfileNameProperty.put("description", "");
        properties.put("ProfileName", ProfileNameProperty);
        Map<String, Object> TemperatureUnitProperty = new HashMap<>();
        TemperatureUnitProperty.put("type", "string");
        TemperatureUnitProperty.put("required", false);
        TemperatureUnitProperty.put("description", "");
        properties.put("TemperatureUnit", TemperatureUnitProperty);
        Map<String, Object> DistanceUnitProperty = new HashMap<>();
        DistanceUnitProperty.put("type", "string");
        DistanceUnitProperty.put("required", false);
        DistanceUnitProperty.put("description", "");
        properties.put("DistanceUnit", DistanceUnitProperty);
        Map<String, Object> TimezoneProperty = new HashMap<>();
        TimezoneProperty.put("type", "string");
        TimezoneProperty.put("required", false);
        TimezoneProperty.put("description", "");
        properties.put("Timezone", TimezoneProperty);
        Map<String, Object> AddressProperty = new HashMap<>();
        AddressProperty.put("type", "string");
        AddressProperty.put("required", false);
        AddressProperty.put("description", "");
        properties.put("Address", AddressProperty);
        Map<String, Object> SetupModeDisabledProperty = new HashMap<>();
        SetupModeDisabledProperty.put("type", "string");
        SetupModeDisabledProperty.put("required", false);
        SetupModeDisabledProperty.put("description", "");
        properties.put("SetupModeDisabled", SetupModeDisabledProperty);
        Map<String, Object> LocaleProperty = new HashMap<>();
        LocaleProperty.put("type", "string");
        LocaleProperty.put("required", false);
        LocaleProperty.put("description", "");
        properties.put("Locale", LocaleProperty);
        Map<String, Object> MaxVolumeLimitProperty = new HashMap<>();
        MaxVolumeLimitProperty.put("type", "string");
        MaxVolumeLimitProperty.put("required", false);
        MaxVolumeLimitProperty.put("description", "");
        properties.put("MaxVolumeLimit", MaxVolumeLimitProperty);
        Map<String, Object> MeetingRoomConfigurationProperty = new HashMap<>();
        MeetingRoomConfigurationProperty.put("type", "string");
        MeetingRoomConfigurationProperty.put("required", false);
        MeetingRoomConfigurationProperty.put("description", "");
        properties.put("MeetingRoomConfiguration", MeetingRoomConfigurationProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_alexa_for_business_update_profile",
            "Updates an existing room profile by room profile ARN.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Alexa_For_Business_Update_ProfileHandler(config));
    }
    
}