package MyApp::Schema::Types;
use strict;
use warnings;
use utf8;

use parent "Exporter";

our @EXPORT = qw/
    PK_INTEGER
    VARCHAR
    INTEGER
/;

sub PK_INTEGER {
    return {
        data_type         => "INTEGER",
        is_nullable       => 0,
        is_auto_increment => 1,
        extra             => { unsigned => 1, },
        @_,
    };
}

sub VARCHAR {
    return {
        data_type   => "VARCHAR",
        size        => 191,
        is_nullable => 0,
    };
}

sub INTEGER {
    return {
        data_type   => "INTEGER",
        is_nullable => 0,
        @_,
    };
}

1;
