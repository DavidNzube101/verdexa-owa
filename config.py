import os

class Config:
    
    SECRET_KEY = os.environ.get('SECRET_KEY') or '737373711hhh'
    
    
    DUNE_API_KEY = os.environ.get('DUNE_API_KEY') or '0tef04bg4lx9drtbjwj5jxlmzgfvep0d'
    DUNE_API_BASE_URL = 'https://api.dune.com/api/v1'
    
    
    CACHE_TIMEOUT = 3600 