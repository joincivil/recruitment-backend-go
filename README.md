# Civil Recruitment Code Test - Golang

Thank you for your interest and taking the time to do our recruitment test. This test is specifically aimed to assess your code skills and proficiency in Golang.

## Time limits

Aim to spend around 3-4 hours on this and return it in 2-3 days. However, we know that
people have busy lives and can't work on these things right away, so feel free to let your contact know when you would be able to submit it. We are more than happy to accommodate.


## The problem

On the Civil Registry, newsrooms create charters that define the mission, ownership structure, revenue model and potential conflicts of interest. This charter can be updated for many reasons, but it is important for the Civil community to understand what has changed and whether these changes affect the status of the newsroom on the Registry.

Your assignment is to build a tool that retrieves the charter revisions data for newsrooms, determine the differences between them and cache those differences. This cached data then needs to be returned via REST to be consumed by a potential client UI. Some components of this project have already been stubbed out and implemented, but it is up to you to complete the work.

To break it down, the service must:

1. Retrieve charter revision data from a set of Civil newsrooms.
2. Determine what has changed between these revisions and serialize this diff data.
3. Cache the diff data to a store.
4. Once the data is persisted, expose the cached diff data via an API.

## How to run

We have included a `main.go` in `cmd/difftool` and in `cmd/api`.

## Requirements

Anything outside these requirements is up to you to design and implement.

* Please make use of the given scaffolding and stubs, but feel free to add any directories, files and additional libraries as needed. Use your favorite dependency manager to add libraries to the repository.
* The charter revision data contains several fields which all need to be diff-ed.
* Diffs only need to be calculated one way from older revision to newer revision.
* A diff between two revisions must determine:
	* The line/column locations of updates between the revisions within the text.
	* The text/character additions, subtractions and changes for those updates.
* Once there is diff data, it must be persisted in some way.
* The API requirements is as follows:
    * It takes 3 parameters used to retrieve the appropriate diff data.
        * Feed name
        * Revision A - the older revision
        * Revision B - the newer revision
    * The data from the API must come from the cache data.
    * The diff data output must be a human readable JSON payload.
    * The diff data output must match the diff requirements above.

## How to submit your work

1. Feel free to add any notes or things you would like us to know about in `NOTES.txt`
2. Zip or tar this entire directory into a file named `<your name>-<role applied for>.{zip|tar}`.
3. Email the file to `codetest@joincivil.com` with the subject `<your name> - <role applied for>`.
