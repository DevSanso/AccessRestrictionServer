package com.github.DevSanso.accessRestrictionGateway.dbc.entity;


import javax.persistence.*;

@Entity
@Table(name = "container")
public class ContainerEntity {
    @Column
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    public int id;
    @Column(name="process_name")
    public String processName;
    @Column
    public int level;
}
