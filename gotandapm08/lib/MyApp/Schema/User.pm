package MyApp::Schema::User;
use 5.20.1;
use warnings;
use utf8;

use parent qw/DBIx::Class::Core/;

__PACKAGE__->table("user");

=pod BEFORE
__PACKAGE__->add_columns(
    id => {
        data_type         => "INTEGER",
        is_nullable       => 0,
        is_auto_increment => 1,
        extra             => { unsigned => 1, },
    },
    name => {
        data_type   => "VARCHAR",
        size        => 191,
        is_nullable => 0,
    },
    level => {
        data_type   => "INTEGER",
        is_nullable => 0,
    },
);
=cut

# AFTER
use MyApp::Schema::Types;

__PACKAGE__->add_columns(
    id                         => PK_INTEGER(),
    name                       => VARCHAR(),
    level                      => INTEGER(),
    main_card_id               => INTEGER(),
    guild_id                   => INTEGER(),
    newest_user_achievement_id => INTEGER(),
);


__PACKAGE__->add_unique_constraint([qw/name/]);

__PACKAGE__->has_one(
    user_main_card               => "Kuwata::Schema::Result::UserCard",
    {"foreign.id"                => "self.main_card_id"},
    {"is_foreign_key_constraint" => 0},
);

__PACKAGE__->might_have(
    newest_user_achievement      => "Kuwata::Schema::Result::UserAchievement",
    {"foreign.id"                => "self.newest_user_achievement_id"},
    {"is_foreign_key_constraint" => 0},
);

__PACKAGE__->has_many(
    user_items => "Kuwata::Schema::Result::UserItem",
    {"foreign.user_id"           => "self.id"},
    {"is_foreign_key_constraint" => 0},
);

__PACKAGE__->belongs_to(
    guild                        => "Kuwata::Schema::Result::Guild",
    {"foreign.id"                => "self.guild_id"},
    {"is_foreign_key_constraint" => 0},
);


1;
