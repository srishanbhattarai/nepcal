# cross

The `cross` binary is meant for internal uses to verify the correctness of the date conversions and to make sure there
are no regressions in correctness. 

During the course of development of `nepcal`, there were several test cases added to make sure that things converted correctly.
Several bugs and edge cases have been found/reported, and subsequently squashed. However, the guarantee that *all* such cases
have been resolved could not be provided. This was further compounded by the fact that different sources on this matter present conflicting information
for several edge cases. 

The project now uses `nepcal.com` as the source of truth for the data that it uses and this binary checks every possible date
against the data dump from that website (`reference.json`).
