#!/usr/bin/env perl
use Mojolicious::Lite -signatures;

my @messages = (
    'hogehoge',
);

plugin "AutoReload";

get "/" => sub ($c) {
    $c->stash(messages => \@messages);
    $c->render(template => "index");
};

post "/messages" => sub ($c) {
    my $message = $c->param("message");
    push @messages, $message;
    $c->stash(messages => \@messages);
    $c->redirect_to("/");
};

get "/list" => sub ($c) {
    $c->render(template => "list");
};

app->renderer->cache->max_keys(0);
app->start;
