# Importing flask module in the project is mandatory
# An object of Flask class is our WSGI application.
from flask import Flask, jsonify
from models import UserInfoRepository

# Flask constructor takes the name of
# current module (__name__) as argument.
app = Flask(__name__)

# The route() function of the Flask class is a decorator,
# which tells the application which URL should call
# the associated function.
@app.route("/")
# ‘/’ URL is bound with hello_world() function.
def hello_world():
    return "Hello World"


@app.route("/hello/<user_name>")
def say_hello(user_name: str = None):
    if user_name:
        return f"Hello, {user_name}"
    else:
        return 'Please provide a name "/hello/\{user_name\}"'


@app.route("/api/getUserInfo/<int:limit>")
def get_all_user_info(limit: int):
    user_repo: UserInfoRepository = UserInfoRepository()
    return jsonify(user_repo.get_all_users(limit))


# main driver function
if __name__ == "__main__":
    # run() method of Flask class runs the application
    # on the local development server.
    app.run(debug=True)
