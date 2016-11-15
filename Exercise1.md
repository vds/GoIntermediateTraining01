# tedfeed go package
We will create an application to download the newst videos from the popular website TED.com.
We are in no way associated with TED.com, but we love it! :)
To do this we will use the [atom feed](https://en.wikipedia.org/wiki/Atom_(standard)) the website provides.

### Setup go package
Let's setup the go package directory

    $> mkdir -p $GOPATH/src/tedfeed/
    $> mkdir -p $GOPATH/src/tedfeed/cmd


### Setup application home folder
Check the existence of the folder "~/tedfeed", "~/tedfeed/videos",
"~/tedfeed/thumbnails" if they don't exists, create them.
Hint: os.Stat, os.Mkdir, os.MkdirAll

Create main package and main function

Hint: is is common for go application to have modules for binaries in the cmd folder

### Download the Ted.com atom feed
Download the atom feed from ted.com: "https://www.ted.com/talks/atom" and print the size of the
feed to the screen.

Hint: http.Get, ioutil.ReadAll

Hint: don't forget to close the body of the response, defer is your friend.
