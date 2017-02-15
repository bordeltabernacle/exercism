local Hamming = {}

function Hamming.compute(strandA, strandB)
	local distance = 0
	for i = 1, string.len(strandA) do
		if string.sub(strandA, i, i) ~= string.sub(strandB, i, i) then
			distance = distance + 1
		end
	end
	return distance
end

return Hamming
