'use strict';

function Hamming() {}

Hamming.prototype.compute = function (strandA, strandB) {

    if (strandA.length != strandB.length) {
        throw new Error('DNA strands must be of equal length.');
    }

    var difference = 0;

    for (var i = 0; i < strandA.length; i++) {
        if (strandA[i] != strandB[i]) {
            difference++;
        };
    };

    return difference;
}

module.exports = Hamming;
