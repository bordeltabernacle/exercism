'use strict';

function Hamming() {}

Hamming.prototype.compute = function (a, b) {

  if (a.length != b.length) {
    throw new Error('DNA strands must be of equal length.');
  }
  var distance = 0;
  for (var i = 0; i < a.length; i++) {
    if (a[i] != b[i]) {
      distance++;
    };
  };
  return distance;
}

module.exports = Hamming;
