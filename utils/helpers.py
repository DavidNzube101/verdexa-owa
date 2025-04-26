import json
from datetime import datetime, timedelta

def generate_date_range(days=30):
    """Generate a date range for the last n days"""
    end_date = datetime.now()
    start_date = end_date - timedelta(days=days)
    
    date_range = []
    current_date = start_date
    
    while current_date <= end_date:
        date_range.append(current_date.strftime('%Y-%m-%d'))
        current_date += timedelta(days=1)
    
    return date_range

def format_dune_response(response):
    """Format a Dune API response for easier consumption"""
    try:
        if 'error' in response:
            return response
        
        result = response.get('result', {})
        rows = result.get('rows', [])
        
        return rows
    except Exception as e:
        return {'error': str(e)}