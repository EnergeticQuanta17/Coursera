from flask import Flask
app=Flask("My first application")

@app.route('/')       # here we have not specified which function to use --> by default it is PUT
def hello():
    return 'Hello World!'

if __name__=="__main__":
    app.run(debug=True)
