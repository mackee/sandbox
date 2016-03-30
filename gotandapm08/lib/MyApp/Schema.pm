package MyApp::Schema;
use 5.20.1;
use warnings;
use utf8;

use parent qw/DBIx::Class::Schema/;

__PACKAGE__->load_namespaces();

1;
