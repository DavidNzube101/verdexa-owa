from services.dune_service import DuneService

class DataService:
    def __init__(self):
        self.dune_service = DuneService()
    
    def get_transaction_flow_data(self, token_address=None):
        """Get transaction flow data for a token"""
        # In a real implementation, this would use the token_address to filter data
        # For now, we just pass a query ID to get dummy data
        return self.dune_service.execute_query("transaction_flow", {"token_address": token_address})
    
    def get_anomaly_data(self, token_address=None, days=14):
        """Get anomaly detection data for a token"""
        params = {"token_address": token_address, "days": days}
        return self.dune_service.execute_query("anomaly_detection", params)
    
    def get_ownership_data(self, token_address=None):
        """Get ownership concentration data for a token"""
        return self.dune_service.execute_query("ownership_concentration", {"token_address": token_address})
    
    def get_sell_off_data(self, token_address=None, days=7):
        """Get sell-off pattern data for a token"""
        params = {"token_address": token_address, "days": days}
        return self.dune_service.execute_query("sell_off_patterns", params)
    
    def get_volume_bracket_data(self, launchpad=None, days=30):
        """Get volume bracket distribution data"""
        params = {"launchpad": launchpad, "days": days}
        return self.dune_service.execute_query("volume_brackets", params)
    
    def get_bot_volume_data(self, token_address=None):
        """Get bot volume detection data for a token"""
        return self.dune_service.execute_query("bot_volume", {"token_address": token_address})
    
    def get_post_rug_data(self, token_address=None):
        """Get post-rug indicators data for a token"""
        return self.dune_service.execute_query("post_rug_indicators", {"token_address": token_address})
    
    def get_wallet_clustering_data(self, token_address=None):
        """Get wallet clustering data for a token"""
        return self.dune_service.execute_query("wallet_clustering", {"token_address": token_address})
    
    def get_dashboard_summary(self):
        """Get dashboard summary data"""
        return self.dune_service.execute_query("dashboard_summary")