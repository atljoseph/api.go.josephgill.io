# logger

This package is a small wrapper around logrus.

It basicaly exports flavors of contextualized logrus *Entry.

The caller is responsible for executing Info(), Warn(), Error(), etc on the resulting *Entry.

Allows the caller to append fields in a structured way that makes sense to the caller.
