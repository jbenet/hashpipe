mdbin = bin/markdown
mdurl = github.com/russross/blackfriday-tool

all: index.html

$(mdbin):
	go get -d $(mdurl)
	go build -o $(mdbin) $(mdurl)

index.html: readme.md footer.html header.html $(mdbin)
	cat header.html >$@
	$(mdbin) -page=false <readme.md >>$@
	cat footer.html >>$@

clean:
	rm index.html

.PHONY: all clean
