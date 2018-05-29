class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/srishanbhattarai/nepcal/archive/v0.3.0.tar.gz"
  head "https://github.com/srishanbhattarai/nepcal.git"
  version "0.3.0"
  sha256 "42275870551a79aff91235a91c72ef89d3d4fb3ddf48a403bfec1d1135f1f14f"
  depends_on "make" => :build

  def install
    bin.install "nepcal"
  end
end
