/**
 * Main MCP Server - Handles MCP JSON-RPC protocol
 */

import java.io.*;
import java.util.*;
import java.util.function.Function;
import java.util.concurrent.CompletableFuture;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.fasterxml.jackson.databind.node.ArrayNode;

// Import Post_X_Amz_Target_Alexa_For_Business_Associate_Skill_With_Skill_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Untag_ResourceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Associate_Skill_With_UsersMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Skill_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_Conference_ProviderMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_DevicesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_DeviceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Disassociate_Device_From_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Business_Report_SchedulesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_UserMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Business_Report_ScheduleMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_DeviceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_Gateway_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Room_Skill_ParameterMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Conference_ProviderMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Associate_Device_With_Network_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_GatewaysMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Approve_SkillMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Forget_Smart_Home_AppliancesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_Network_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_ProfilesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Gateway_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_GatewayMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Conference_PreferenceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_ContactMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_ContactMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_TagsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_UserMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Resolve_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Skills_Store_Skills_By_CategoryMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_GatewayMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_Conference_ProviderMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_ContactMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Invitation_ConfigurationMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Put_Invitation_ConfigurationMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_Gateway_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Skills_Store_CategoriesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Room_Skill_ParameterMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Device_Usage_DataMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_Business_Report_ScheduleMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Put_Skill_AuthorizationMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_Skill_GroupsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_RoomsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Skill_AuthorizationMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Address_BookMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Device_EventsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_ContactsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_Skill_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_Group_From_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_UsersMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Associate_Contact_With_Address_BookMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_ContactMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_From_Skill_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Put_Conference_PreferenceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Tag_ResourceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Gateway_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Disassociate_Contact_From_Address_BookMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Send_InvitationMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Register_Avs_DeviceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Smart_Home_AppliancesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Reject_SkillMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_Address_BookMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_DeviceMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_SkillsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Update_Address_BookMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Create_Skill_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Gateway_GroupsMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Start_Smart_Home_Appliance_DiscoveryMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_List_Conference_ProvidersMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Skill_GroupMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Revoke_InvitationMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_Address_BooksMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Conference_ProviderMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Associate_Skill_Group_With_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Start_Device_SyncMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Search_Network_ProfilesMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_From_UsersMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Get_Network_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Address_BookMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Associate_Device_With_RoomMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Delete_Network_ProfileMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Put_Room_Skill_ParameterMCPTool - included in same package
// Import Post_X_Amz_Target_Alexa_For_Business_Send_AnnouncementMCPTool - included in same package

public class MCPServer {
    
    // Common classes that all tool classes use
    public static class APIConfig {
        private final String baseUrl;
        private final String apiKey;
        
        public APIConfig(String baseUrl, String apiKey) {
            this.baseUrl = baseUrl;
            this.apiKey = apiKey;
        }
        
        public String getBaseUrl() { return baseUrl; }
        public String getApiKey() { return apiKey; }
    }
    
    public static class MCPRequest {
        private Map<String, Object> params;
        
        public Map<String, Object> getParams() { return params; }
        public void setParams(Map<String, Object> params) { this.params = params; }
        
        @SuppressWarnings("unchecked")
        public Map<String, Object> getArguments() {
            if (params != null && params.containsKey("arguments")) {
                return (Map<String, Object>) params.get("arguments");
            }
            return new HashMap<>();
        }
    }
    
    public static class MCPToolResult {
        private final String content;
        private final boolean isError;
        
        public MCPToolResult(String content, boolean isError) {
            this.content = content;
            this.isError = isError;
        }
        
        public MCPToolResult(String content) {
            this(content, false);
        }
        
        public String getContent() { return content; }
        public boolean isError() { return isError; }
    }
    
    public static class ToolDefinition {
        private final String name;
        private final String description;
        private final Map<String, Object> parameters;
        
        public ToolDefinition(String name, String description, Map<String, Object> parameters) {
            this.name = name;
            this.description = description;
            this.parameters = parameters;
        }
        
        public String getName() { return name; }
        public String getDescription() { return description; }
        public Map<String, Object> getParameters() { return parameters; }
    }
    
    public static class Tool {
        private final ToolDefinition definition;
        private final Function<MCPRequest, MCPToolResult> handler;
        
        public Tool(ToolDefinition definition, Function<MCPRequest, MCPToolResult> handler) {
            this.definition = definition;
            this.handler = handler;
        }
        
        public ToolDefinition getDefinition() { return definition; }
        public Function<MCPRequest, MCPToolResult> getHandler() { return handler; }
    }
    
    private static final ObjectMapper mapper = new ObjectMapper();
    private static List<Tool> tools = new ArrayList<>();
    private static APIConfig config;
    
    public static void main(String[] args) {
        try {
            // Initialize configuration
            String baseUrl = System.getenv("API_BASE_URL");
            String apiKey = System.getenv("API_KEY");
            
            if (baseUrl == null || apiKey == null) {
                System.err.println("Error: Please set API_BASE_URL and API_KEY environment variables");
                System.exit(1);
            }
            
            config = new APIConfig(baseUrl, apiKey);
            
            // Register all tools
            tools = Arrays.asList(
            Post_X_Amz_Target_Alexa_For_Business_Associate_Skill_With_Skill_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Associate_Skill_With_Skill_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Untag_ResourceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Untag_ResourceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Associate_Skill_With_UsersMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Associate_Skill_With_UsersTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Skill_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Skill_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_Conference_ProviderMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_Conference_ProviderTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_DevicesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_DevicesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_DeviceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_DeviceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Disassociate_Device_From_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Disassociate_Device_From_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Business_Report_SchedulesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Business_Report_SchedulesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_UserMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_UserTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Business_Report_ScheduleMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Business_Report_ScheduleTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_DeviceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_DeviceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_Gateway_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_Gateway_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Room_Skill_ParameterMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Room_Skill_ParameterTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Conference_ProviderMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Conference_ProviderTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Associate_Device_With_Network_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Associate_Device_With_Network_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_GatewaysMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_GatewaysTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Approve_SkillMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Approve_SkillTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Forget_Smart_Home_AppliancesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Forget_Smart_Home_AppliancesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_Network_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_Network_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_ProfilesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_ProfilesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Gateway_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Gateway_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_GatewayMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_GatewayTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Conference_PreferenceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Conference_PreferenceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_ContactMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_ContactTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_ContactMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_ContactTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_TagsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_TagsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_UserMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_UserTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Resolve_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Resolve_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Skills_Store_Skills_By_CategoryMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Skills_Store_Skills_By_CategoryTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_GatewayMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_GatewayTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_Network_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_Conference_ProviderMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_Conference_ProviderTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_ContactMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_ContactTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Invitation_ConfigurationMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Invitation_ConfigurationTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Put_Invitation_ConfigurationMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Put_Invitation_ConfigurationTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_Gateway_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_Gateway_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Skills_Store_CategoriesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Skills_Store_CategoriesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Room_Skill_ParameterMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Room_Skill_ParameterTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Device_Usage_DataMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Device_Usage_DataTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_Business_Report_ScheduleMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_Business_Report_ScheduleTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Put_Skill_AuthorizationMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Put_Skill_AuthorizationTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_Skill_GroupsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_Skill_GroupsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_RoomsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_RoomsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Skill_AuthorizationMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Skill_AuthorizationTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Address_BookMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Address_BookTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Device_EventsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Device_EventsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_Business_Report_ScheduleTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_ContactsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_ContactsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_Skill_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_Skill_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_Group_From_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_Group_From_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_UsersMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_UsersTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Associate_Contact_With_Address_BookMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Associate_Contact_With_Address_BookTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_ContactMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_ContactTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_From_Skill_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_From_Skill_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Put_Conference_PreferenceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Put_Conference_PreferenceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Tag_ResourceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Tag_ResourceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Gateway_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Gateway_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Disassociate_Contact_From_Address_BookMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Disassociate_Contact_From_Address_BookTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Send_InvitationMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Send_InvitationTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Register_Avs_DeviceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Register_Avs_DeviceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Smart_Home_AppliancesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Smart_Home_AppliancesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Reject_SkillMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Reject_SkillTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_Address_BookMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_Address_BookTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_DeviceMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_DeviceTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_SkillsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_SkillsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Update_Address_BookMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Update_Address_BookTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Create_Skill_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Create_Skill_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Gateway_GroupsMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Gateway_GroupsTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Start_Smart_Home_Appliance_DiscoveryMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Start_Smart_Home_Appliance_DiscoveryTool(config),
            Post_X_Amz_Target_Alexa_For_Business_List_Conference_ProvidersMCPTool.createPost_X_Amz_Target_Alexa_For_Business_List_Conference_ProvidersTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Skill_GroupMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Skill_GroupTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Revoke_InvitationMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Revoke_InvitationTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_Address_BooksMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_Address_BooksTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Conference_ProviderMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Conference_ProviderTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Associate_Skill_Group_With_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Associate_Skill_Group_With_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Start_Device_SyncMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Start_Device_SyncTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Search_Network_ProfilesMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Search_Network_ProfilesTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_From_UsersMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Disassociate_Skill_From_UsersTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Get_Network_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Get_Network_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Address_BookMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Address_BookTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Associate_Device_With_RoomMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Associate_Device_With_RoomTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Delete_Network_ProfileMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Delete_Network_ProfileTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Put_Room_Skill_ParameterMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Put_Room_Skill_ParameterTool(config),
            Post_X_Amz_Target_Alexa_For_Business_Send_AnnouncementMCPTool.createPost_X_Amz_Target_Alexa_For_Business_Send_AnnouncementTool(config)
            );
            
            System.err.println("MCP Server started with " + tools.size() + " tools");
            
            // Start JSON-RPC server
            runServer();
            
        } catch (Exception e) {
            System.err.println("Failed to start MCP server: " + e.getMessage());
            e.printStackTrace();
            System.exit(1);
        }
    }
    
    private static void runServer() throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String line;
        
        while ((line = reader.readLine()) != null) {
            JsonNode request = null;
            try {
                request = mapper.readTree(line);
                JsonNode response = handleRequest(request);
                
                if (response != null) {
                    System.out.println(mapper.writeValueAsString(response));
                }
                
            } catch (Exception e) {
                // Send error response
                ObjectNode errorResponse = mapper.createObjectNode();
                errorResponse.put("jsonrpc", "2.0");
                if (request != null && request.has("id")) {
                    errorResponse.put("id", request.get("id").asText());
                } else {
                    errorResponse.putNull("id");
                }
                
                ObjectNode error = mapper.createObjectNode();
                error.put("code", -32603);
                error.put("message", "Internal error: " + e.getMessage());
                errorResponse.set("error", error);
                
                System.out.println(mapper.writeValueAsString(errorResponse));
            }
        }
    }
    
    private static JsonNode handleRequest(JsonNode request) {
        if (!request.has("method")) {
            return createErrorResponse(request, -32600, "Invalid Request");
        }
        
        String method = request.get("method").asText();
        JsonNode params = request.get("params");
        String id = request.has("id") ? request.get("id").asText() : null;
        
        switch (method) {
            case "initialize":
                return handleInitialize(id);
            case "tools/list":
                return handleToolsList(id);
            case "tools/call":
                return handleToolsCall(id, params);
            default:
                return createErrorResponse(request, -32601, "Method not found");
        }
    }
    
    private static JsonNode handleInitialize(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        result.put("protocolVersion", "2024-11-05");
        
        ObjectNode capabilities = mapper.createObjectNode();
        ObjectNode toolsCapability = mapper.createObjectNode();
        toolsCapability.put("listChanged", false);
        capabilities.set("tools", toolsCapability);
        result.set("capabilities", capabilities);
        
        ObjectNode serverInfo = mapper.createObjectNode();
        serverInfo.put("name", "Opsera MCP Server");
        serverInfo.put("version", "1.0.0");
        result.set("serverInfo", serverInfo);
        
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsList(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        ArrayNode toolsArray = mapper.createArrayNode();
        
        for (Tool tool : tools) {
            ObjectNode toolDef = mapper.createObjectNode();
            toolDef.put("name", tool.getDefinition().getName());
            toolDef.put("description", tool.getDefinition().getDescription());
            
            // Convert parameters to JSON
            ObjectNode inputSchema = mapper.valueToTree(tool.getDefinition().getParameters());
            toolDef.set("inputSchema", inputSchema);
            
            toolsArray.add(toolDef);
        }
        
        result.set("tools", toolsArray);
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsCall(String id, JsonNode params) {
        if (!params.has("name")) {
            return createErrorResponse(null, -32602, "Invalid params: missing 'name'");
        }
        
        String toolName = params.get("name").asText();
        JsonNode arguments = params.has("arguments") ? params.get("arguments") : mapper.createObjectNode();
        
        // Find the tool
        Tool tool = null;
        for (Tool t : tools) {
            if (t.getDefinition().getName().equals(toolName)) {
                tool = t;
                break;
            }
        }
        
        if (tool == null) {
            return createErrorResponse(null, -32602, "Tool not found: " + toolName);
        }
        
        try {
            // Convert arguments to Map
            Map<String, Object> argsMap = mapper.convertValue(arguments, Map.class);
            
            // Create MCP request
            MCPRequest mcpRequest = new MCPRequest();
            Map<String, Object> requestParams = new HashMap<>();
            requestParams.put("arguments", argsMap);
            mcpRequest.setParams(requestParams);
            
            // Execute tool
            MCPToolResult result = tool.getHandler().apply(mcpRequest);
            
            // Create response
            ObjectNode response = mapper.createObjectNode();
            response.put("jsonrpc", "2.0");
            response.put("id", id);
            
            ObjectNode resultObj = mapper.createObjectNode();
            ArrayNode content = mapper.createArrayNode();
            
            ObjectNode textContent = mapper.createObjectNode();
            textContent.put("type", "text");
            textContent.put("text", result.getContent());
            content.add(textContent);
            
            resultObj.set("content", content);
            resultObj.put("isError", result.isError());
            
            response.set("result", resultObj);
            return response;
            
        } catch (Exception e) {
            return createErrorResponse(null, -32603, "Tool execution failed: " + e.getMessage());
        }
    }
    
    private static JsonNode createErrorResponse(JsonNode request, int code, String message) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", request != null && request.has("id") ? request.get("id").asText() : null);
        
        ObjectNode error = mapper.createObjectNode();
        error.put("code", code);
        error.put("message", message);
        response.set("error", error);
        
        return response;
    }
}