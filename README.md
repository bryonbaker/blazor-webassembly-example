# Steps

## Setup

`dotnet new blasor-wasm`

`dotnet add package Outrage.Patternfly`

`dotnet run --urls "http://*:8080"`

Browse to localhost to test app is running

## Remove all the boostrap shit

### Index.html
In wwwroot... 

Replace `<link href="css/bootstrap/bootstrap.min.css" rel="stylesheet" />` with:

```
<link href="_content/Outrage.Patternfly/css/patternfly.base.css" rel="stylesheet" />
<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" rel="stylesheet" />
```

Delete uneeded folders: `./css/bootstrap` and `./open-iconic`

Running app will now be missing all thr styling shit

## Incorporate Patternfly

From Outrage, click the Patternfly url

Paste the content into MainLayout.razor (leave the inerits line at the top)

```
<PatternflyPage FillMainContent="true">
    <Logo>
        <img class="pf-c-brand pf-u-display-block pf-u-display-none-on-md" style="height: 64px;"src="/images/icon_only.svg" alt="Patternfly">
        <img class="pf-c-brand pf-u-display-none pf-u-display-block-on-md" style="height: 64px;" src="/images/icon.svg" alt="Patternfly">
    </Logo>
    <Header>
        <PatternflyNavigation Horizontal="true">
            <PatternflyNavigationGroup>
                <PatternflyNavigationLink Href="" Match="NavLinkMatch.All">Home</PatternflyNavigationLink>
            </PatternflyNavigationGroup>
        </PatternflyNavigation>
    </Header>
    <Sections>
        <PatternflyPageSection Fill="true" Color="SectionColor.normal">
            @Body
        </PatternflyPageSection>
        <PatternflyPageSection Color="SectionColor.dark_200">
            Outrage Patternfly is Copyright (C) Webefinity Pty. Ltd. 2022.
        </PatternflyPageSection>
    </Sections>
    <Sidebar>
        <PatternflyNavigation>
            <PatternflyNavigationGroup>
                <PatternflyNavigationLink Href="getting-started" Icon="play">Getting Started</PatternflyNavigationLink>
                <PatternflyNavigationLink Href="patternfly-page" Icon="newspaper">Patternfly Page</PatternflyNavigationLink>
                ...
            </PatternflyNavigationGroup>
        </PatternflyNavigation>
    </Sidebar>
</PatternflyPage>
```

Add these line to the top of the file:

```
@using Outrage.Patternfly.Components.Navigation
@using Outrage.Patternfly.Components.Page
```