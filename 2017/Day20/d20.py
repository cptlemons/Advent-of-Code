import re
import sys

with open ('2017/Day20/d20input.txt') as i:
    particles = []
    for line in i:
        particle = []
        vals = re.split('<|>',line.strip())[1::2]
        for val in vals:
            particle.append([int(v) for v in val.split(',')])
        particles.append(particle)

def part1(particles):
    lowest_p = [1e100]
    for i,p in enumerate(particles):
        total_a = sum([abs(v) for v in p[2]])
        if total_a > lowest_p[0]:
            continue
        elif total_a == lowest_p[0]:
            lowest_p.append(i)
            continue
        else:
            lowest_p = [total_a, i]
    return lowest_p

print(part1(particles))

def part2(particles):
    for j in range(100):
        locs = {}
        collisions = set()
        for idx, particle in enumerate(particles):
            x_pos = particle[0][0] + particle[1][0] + particle[2][0]
            y_pos = particle[0][1] + particle[1][1] + particle[2][1]
            z_pos = particle[0][2] + particle[1][2] + particle[2][2]
            if (x_pos, y_pos, z_pos) in locs:
                print("Collision at {} at time {}".format((x_pos,y_pos,z_pos),j))
                locs[(x_pos, y_pos, z_pos)] += [idx]
                collisions.add((x_pos, y_pos, z_pos))
            else:
                locs[(x_pos, y_pos, z_pos)] = [idx]
            particles[idx][0] = [x_pos, y_pos, z_pos]
            particles[idx][1] = [particle[1][0] + particle[2][0], particle[1][1] + particle[2][1], particle[1][2] + particle[2][2]]
        to_remove = []
        for col in collisions:
            for p in locs[col]:
                to_remove += [p]
        for p in reversed(sorted(to_remove)):
            del particles[p]
        if j % 10 == 0:
            print(len(particles))

part2(particles)