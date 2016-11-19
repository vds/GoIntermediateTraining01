# tedfeed go package
We'll create an application to download the newst videos from the popular website TED.com.
We are in no way associated with TED.com, but we love it! :)
To do this we will use the [atom feed](https://en.wikipedia.org/wiki/Atom_(standard)) the website provides.
We'll learn how to interact with the OS environment, how to deal with HTTP requests and buffers.

### Setup go package
Let's setup the go package directory

    $> mkdir -p $GOPATH/src/tedfeed/
    $> mkdir -p $GOPATH/src/tedfeed/cmd


### Setup application home folder
Create main package and main function

Hint: it is common for go application to have modules for binaries in the cmd folder

Check the existence of the folder "~/tedfeed", "~/tedfeed/videos",
"~/tedfeed/thumbnails" if they don't exists, create them.

Hint: **user.Current**, **os.Getenv** **os.Stat**, **os.Mkdir**, **os.MkdirAll**


### Download the Ted.com atom feed
Download the atom feed from ted.com: "https://www.ted.com/talks/atom" and print the size of the
feed to the screen.

Hint: **http.Get**, **ioutil.ReadAll**

Hint: don't forget to close the body of the response, defer is your friend.
