from flask import Flask, render_template, request
from dht11_example import getResult
app = Flask(__name__)

@app.route('/humidity')
def onReqHumidity():
	result = getResult()
	if result.is_valid():
		print("Humid:", result.humidity)

@app.route('/temperature')
def onReqTemperature():
	result = getResult()
	if result.is_valid():
		print("Temp:", result.temperature)


if __name__ == "__main__":
   app.run(host='0.0.0.0', port=80, debug=True)
