from flask import Flask, jsonify
import os

app = Flask(__name__)

@app.route('/')
def home():
    return "Welcome to the Flask App!"

@app.route('/health')
def health_check():
    return jsonify({"status": "OK", "message": "Application is running"}), 200

if __name__ == '__main__':
    port = int(os.environ.get("PORT", 5000))
    app.run(host='0.0.0.0', port=port)