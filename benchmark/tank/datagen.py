#!/usr/bin/python                                                               
# -*- coding: utf-8 -*-                                                         
from ammo  import AmmoGenerator

HOST="192.168.99.100"
URI='/test-service'
FILE="ammo.txt"
COUNT=100

def main():
    gen = AmmoGenerator(method="GET")
    with open(FILE, 'w') as expfile:
        for x in xrange(COUNT):
            expfile.write(gen.full_request(
                HOST, None, case=URI
            ))
    print "Created %d ammo in %s" % (COUNT, FILE)

if __name__ == '__main__':
    main()
