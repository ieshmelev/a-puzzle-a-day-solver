package main

type solver struct {
	l  *logger
	cs []coord
}

func (s *solver) solve(f field, pool []set) (field, bool) {
	if len(pool) == 0 {
		return f, true
	}
	for k, pieces := range pool {
		for _, p := range pieces {
			nf, ok := f.put(p)
			if !ok {
				continue
			}
			if !nf.isEmpty(s.cs...) {
				continue
			}
			s.l.Debug(nf)
			npool := make([]set, len(pool))
			copy(npool, pool)
			solution, ok := s.solve(
				nf,
				append(npool[:k], npool[k+1:]...),
			)
			if ok {
				return solution, true
			}
		}
	}
	return nil, false

}
