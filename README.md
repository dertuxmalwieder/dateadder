# dateadder

Fulfills the single need of having a platform-independent tool that lets me, the developer, add days, weeks or even months to today's (or any other) date and see the resulting date.

* Does not support units smaller than a day. Feel free to fork this project and make a `houradder` or something.
* Also does not support substracting days, weeks or even months. Feel free to fork this project and make a `datesubstracter` or something.

Posted here for your convenience.

## Examples

    % ./dateadder "today in a week"
    % ./dateadder "2021/12/31 plus four days"

Et cetera, et cetera.

## URLs

Upstream Fossil repository:
https://code.rosaelefanten.org/dateadder

Git clone:
https://github.com/dertuxmalwieder/dateadder

Note that I will only actively use the Fossil one. Please contact me if you want to contribute.

## Building and installation

    % fossil clone https://code.rosaelefanten.org/dateadder dateadder.fossil
    % mkdir dateadder ; cd dateadder ; fossil open ../dateadder.fossil
    % cd src
    % go build

## Support

No.
