local songlines = {
  {'house', 'Jack built.'},
  {'malt', 'lay in '},
  {'rat', 'ate '},
  {'cat', 'killed '},
  {'dog', 'worried '},
  {'cow with the crumpled horn', 'tossed '},
  {'maiden all forlorn', 'milked '},
  {'man all tattered and torn', 'kissed '},
  {'priest all shaven and shorn', 'married '},
  {'rooster that crowed in the morn', 'woke '},
  {'farmer sowing his corn', 'kept '},
  {'horse and the hound and the horn', 'belonged to '}
}

local house = {}

house.verse = function(n)
  local verse = ''
  for i = 1, n do
    local noun, verb = unpack(songlines[i])
    if i == 1 then that = ' that ' else that = '\n  that ' end
    verse = 'the ' .. noun .. that .. verb .. verse
  end
  return 'This is ' .. verse
end

house.recite = function()
  local song = {}
  for i = 1, #songlines do
    table.insert(song, house.verse(i))
  end
  return table.concat(song, '\n')
end

return house
