package com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers;


import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Entity
@Table(name = "container")
@Builder
@Getter
@NoArgsConstructor
@AllArgsConstructor
public class Container {
    @Column
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    public Long id;
    @Column(name="process_name")
    public String processName;
    @Column
    public int level;
}
