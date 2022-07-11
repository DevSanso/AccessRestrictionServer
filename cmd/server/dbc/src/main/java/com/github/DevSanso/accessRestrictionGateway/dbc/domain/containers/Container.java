package com.github.DevSanso.accessRestrictionGateway.dbc.domain.containers;


import javax.persistence.*;

@Entity
@Table(name = "container")
public class Container {
    @Column
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    public Long id;
    @Column(name="process_name")
    public String processName;
    @Column
    public int level;
}
