local songlines = {
  ' the house that Jack built.',
  ' the malt\nthat lay in',
  ' the rat\nthat ate',
  ' the cat\nthat killed',
  ' the dog\nthat worried',
  ' the cow with the crumpled horn\nthat tossed',
  ' the maiden all forlorn\nthat milked',
  ' the man all tattered and torn\nthat kissed',
  ' the priest all shaven and shorn\nthat married',
  ' the rooster that crowed in the morn\nthat woke',
  ' the farmer sowing his corn\nthat kept',
  ' the horse and the hound and the horn\nthat belonged to',
}

local house = {}

house.verse = function(n)
  local v = ''
  for i = 1, n do
    v = songlines[i] .. v
  end
  return 'This is' .. v
end

house.recite = function()
  local s = ''
  for i = 1, #songlines do
    s = s .. house.verse(i)
    if i ~= #songlines then s = s .. '\n' end
  end
  return s
end

return house
