from flask import Blueprint, jsonify, request
from services.data_service import DataService

api_bp = Blueprint('api', __name__)
data_service = DataService()

@api_bp.route('/transaction-flow', methods=['GET'])
def transaction_flow():
    token_address = request.args.get('token_address')
    data = data_service.get_transaction_flow_data(token_address)
    return jsonify(data)

@api_bp.route('/anomaly-detection', methods=['GET'])
def anomaly_detection():
    token_address = request.args.get('token_address')
    days = request.args.get('days', default=14, type=int)
    data = data_service.get_anomaly_data(token_address, days)
    return jsonify(data)

@api_bp.route('/ownership-concentration', methods=['GET'])
def ownership_concentration():
    token_address = request.args.get('token_address')
    data = data_service.get_ownership_data(token_address)
    return jsonify(data)

@api_bp.route('/sell-off-patterns', methods=['GET'])
def sell_off_patterns():
    token_address = request.args.get('token_address')
    days = request.args.get('days', default=7, type=int)
    data = data_service.get_sell_off_data(token_address, days)
    return jsonify(data)

@api_bp.route('/volume-brackets', methods=['GET'])
def volume_brackets():
    launchpad = request.args.get('launchpad')
    days = request.args.get('days', default=30, type=int)
    data = data_service.get_volume_bracket_data(launchpad, days)
    return jsonify(data)

@api_bp.route('/bot-volume', methods=['GET'])
def bot_volume():
    token_address = request.args.get('token_address')
    data = data_service.get_bot_volume_data(token_address)
    return jsonify(data)

@api_bp.route('/post-rug-indicators', methods=['GET'])
def post_rug_indicators():
    token_address = request.args.get('token_address')
    data = data_service.get_post_rug_data(token_address)
    return jsonify(data)

@api_bp.route('/wallet-clustering', methods=['GET'])
def wallet_clustering():
    token_address = request.args.get('token_address')
    data = data_service.get_wallet_clustering_data(token_address)
    return jsonify(data)

@api_bp.route('/dashboard-summary', methods=['GET'])
def dashboard_summary():
    data = data_service.get_dashboard_summary()
    return jsonify(data)