#!/usr/bin/perl

use strict;
use warnings;
use Math::Trig qw(pi);

# What is the performance impact (relatively) of polymorphism in Perl?

# OOP in Perl is already weird as is. Let's try to do this with Moose.

{
    package Shape;
    use Moose;

    sub area;
}

{ 
    package Square;
    use Moose;

    # extends 'Shape'; 
    # Using inheritance via the 'extends' keyword is wholy unnecessary here
    # Perl doesn't really do type checking. If the 'area()' function exists.
    # It doesn't do type checking of what you put into an array anyway.

    has 'side', is => 'ro', isa => 'Num';

    sub area {
        my $self = shift;
        return $self->side() * $self->side();
    }
}

{ 
    package Rectangle;
    use Moose;

    # extends 'Shape'; # unnecessary

    has 'height', is => 'ro', isa => 'Num';
    has 'width',  is => 'ro', isa => 'Num';

    sub area {
        my $self = shift;
        return $self->height() * $self->width();
    }
}

{
    package Triangle;
    use Moose;
    
    # extends 'Shape'; # unnecessary

    has 'height', is => 'ro', isa => 'Num';
    has 'width',  is => 'ro', isa => 'Num';

    sub area {
        my $self = shift;
        return 0.5 * $self->height() * $self->width();
    }
}

{
    package Circle;
    use Moose;
    use Math::Trig qw(pi);

    # extends 'Shape'; # unnecessary

    has 'radius', is => 'ro', isa => 'Num';

    sub area {
        my $self = shift;
        return  pi * $self->radius() * $self->radius();
    }
}

my @shapeObjects;

@shapeObjects[$_] = Square->new( side => $_ * $_ + 1 )                  for 0 .. 999 ;
@shapeObjects[$_] = Rectangle->new( height => $_ + 1, width => $_ + 1 ) for 1000 .. 1999 ;
@shapeObjects[$_] = Triangle->new( height => $_ + 1, width => $_ + 1 )  for 2000 .. 2999 ;
@shapeObjects[$_] = Circle->new( radius => $_ + 1 )                     for 3000 .. 3999 ;

sub calculateShapeAreasOOPversion {
    my $sum;
    while (my ($index, $shape) = each @shapeObjects) {
        $sum += $shape->area();
    }
    return $sum
}

# Make sure the data is accessed at least once before
print 'sum of all area ', calculateShapeAreasOOPversion(), "\n";

# Array of hashes variation
my @shapesHashes;
@shapesHashes[$_] = { type => 'square',    side => $_ * $_ + 1               } for 0 .. 999 ;
@shapesHashes[$_] = { type => 'rectangle', height => $_ + 1, width => $_ + 1 } for 1000 .. 1999 ;
@shapesHashes[$_] = { type => 'triangle',  height => $_ + 1, width => $_ + 1 } for 2000 .. 2999 ;
@shapesHashes[$_] = { type => 'circle',    radius => $_ + 1                  } for 3000 .. 3999 ;
use Data::Dumper;

sub calculateAreaBasedOnType {
    my $shape = shift;
    if ($shape->{type} eq 'square') {
        return $shape->{side} * $shape->{side};
    } elsif ($shape->{type} eq 'rectangle') {
        return $shape->{height} * $shape->{width};
    } elsif ($shape->{type} eq 'triangle') {
        return 0.5 * $shape->{height} * $shape->{width};
    } elsif ($shape->{type} eq 'circle') {
        return pi * $shape->{radius} * $shape->{radius};
    } else {
        die 'unexpected type';
    }
}

sub calculateShapeAreasHashesVersion {
    my $sum;
    while (my ($index, $shape) = each @shapesHashes) {
        $sum += calculateAreaBasedOnType ($shape);
    }
    return $sum
}

print 'sum of all area ', calculateShapeAreasHashesVersion(), "\n";

# BENCHMARK

use Benchmark qw(timethis timethese cmpthese);

# timethis( -5, 'calculateShapeAreasOOPversion()' );

# timethese( -5, {
#     'calculateShapeAreasOOPversion()' => 'calculateShapeAreasOOPversion()',
#     'calculateShapeAreasHashesVersion()' => 'calculateShapeAreasHashesVersion()'
# });

cmpthese( -10, {
    'calculateShapeAreasOOPversion()' => 'calculateShapeAreasOOPversion()',
    'calculateShapeAreasHashesVersion()' => 'calculateShapeAreasHashesVersion()'
});
