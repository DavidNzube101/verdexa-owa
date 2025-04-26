from flask import Flask, jsonify
from flask_cors import CORS
from api.routes import api_bp
from config import Config

def create_app(config_class=Config):
    app = Flask(__name__)
    app.config.from_object(config_class)
    
    
    CORS(app)
    
    
    app.register_blueprint(api_bp, url_prefix='/api')
    
    @app.route('/health')
    def health_check():
        return jsonify({"status": "ok"})
    
    return app

if __name__ == '__main__':
    app = create_app()
    app.run(debug=True)