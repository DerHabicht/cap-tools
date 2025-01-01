export const enum Uniform {
    MessDress = 'USAF Mess Dress',
    SemiFormal = 'USAF Semi-Formal',
    ServiceDress = 'USAF Service Dress',
    ClassB = 'USAF Class B',
    ABU = 'ABU',
    FDU = 'FDU',
    PTU = 'PTU',
    CorporateSemiFormal = 'Corporate Semi-Formal',
    CorporateServiceDress = 'Corporate Service Dress',
    Aviator = 'Aviator Combination',
    CWU = 'CWU',
    CFU = 'CFU',
    CFDU ='CFDU',
    CivilianAttire = 'Civilian Attire',
}

export function formatUniformList(uniforms: Uniform[], conjunction: 'and' | 'or') {
    const allButLast = uniforms.slice(0, -1);

    return `${allButLast.join(", ")}, ${conjunction} ${uniforms[uniforms.length - 1]}`;
}