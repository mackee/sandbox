#!/usr/bin/env perl
use Mojolicious::Lite -signatures;
use Encode qw/decode_utf8/;
use JSON::XS qw/encode_json/;

use Mojo::WebSocket qw(WS_PING);
use Mojo::IOLoop;

my @messages = (
    { id => 1, message => 'hogehoge' },
);
my $cnt = scalar(@messages);
my %toggles = (
    1 => 0,
);

my %connections;

get "/" => sub ($c) {
    $c->stash(messages => \@messages);
    $c->stash(toggles => \%toggles);
    $c->render(template => "index", layout => "default");
};

post "/messages" => sub ($c) {
    my $message_text = $c->param("message");
    $cnt++;
    my $message = { id => $cnt, message => $message_text };
    push @messages, $message;

    $c->stash(messages => [$message]);
    my $rendered = $c->render_to_string(template => "append_messages");

    for my $connection (values %connections) {
        $connection->send(decode_utf8(encode_json({
            identifier => encode_json({ channel => "Turbo::StreamsChannel", signed_stream_name => "**mysignature**" }),
            message    => $rendered->to_string =~ s/\n//gr,
        })));
    }

    if ($c->req->headers->accept =~ m!text/vnd\.turbo-stream\.html!) {
        $c->finish;
        return;
    }

    $c->redirect_to("/");
};

post "/toggles" => sub ($c) {
    my $checked_values = $c->every_param("checked");
    %toggles = (
        (map { $_ => 0 } keys %toggles),
        (map { $_ => 1 } (defined $checked_values ? @$checked_values : ())),
    );
    if ($c->req->headers->accept =~ m!text/vnd\.turbo-stream\.html!) {
        $c->stash(toggles => \%toggles);
        $c->render(template => "_toggles");
        return;
    }
    $c->redirect_to("/");
};

post "/add_toggle" => sub ($c) {
    my $toggle_cnt = scalar(keys %toggles);
    $toggles{$toggle_cnt+1} = 1;
    if ($c->req->headers->accept =~ m!text/vnd\.turbo-stream\.html!) {
        $c->stash(toggles => \%toggles);
        $c->render(template => "_toggles");
        return;
    }
    $c->redirect_to("/");
};

websocket "/streams" => sub ($c) {
    if ($c->tx->is_websocket) {
        $connections{$c->tx->connection} = $c->tx;
        $c->tx->on("finish" => sub ($c, @ignore) {
            delete $connections{$c->connection};
        });
        $c->on("message" => sub ($c, $msg) {
            warn $msg;
        });
        my $id;
        $id = Mojo::IOLoop->recurring(5 => sub ($loop) {
            if (!defined $c->tx || $c->tx->is_finished) {
                $loop->remove($id);
                return;
            }
            $c->send([1, 0, 0, 0, WS_PING, 'Hello World!']);
        });
    }
};

app->renderer->cache->max_keys(0);
app->start;
